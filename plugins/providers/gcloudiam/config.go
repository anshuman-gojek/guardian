package gcloudiam

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/odpf/guardian/domain"
)

const (
	AccountTypeUser           = "user"
	AccountTypeServiceAccount = "serviceAccount"
	AccountTypeGroup          = "group"
)

type Credentials struct {
	ServiceAccountKey string `mapstructure:"service_account_key" json:"service_account_key" validate:"required,base64"`
	ResourceName      string `mapstructure:"resource_name" json:"resource_name" validate:"startswith=projects/|startswith=organizations/"`
}

func (c *Credentials) Encrypt(encryptor domain.Encryptor) error {
	if c == nil {
		return ErrUnableToEncryptNilCredentials
	}

	encryptedSAKey, err := encryptor.Encrypt(c.ServiceAccountKey)
	if err != nil {
		return err
	}

	c.ServiceAccountKey = encryptedSAKey
	return nil
}

func (c *Credentials) Decrypt(decryptor domain.Decryptor) error {
	if c == nil {
		return ErrUnableToDecryptNilCredentials
	}

	decryptedSAKey, err := decryptor.Decrypt(c.ServiceAccountKey)
	if err != nil {
		return err
	}

	c.ServiceAccountKey = decryptedSAKey
	return nil
}

type Config struct {
	ProviderConfig *domain.ProviderConfig
	valid          bool

	crypto    domain.Crypto
	validator *validator.Validate
}

func NewConfig(pc *domain.ProviderConfig, crypto domain.Crypto) *Config {
	return &Config{
		ProviderConfig: pc,
		validator:      validator.New(),
		crypto:         crypto,
	}
}

func (c *Config) ParseAndValidate() error {
	return c.parseAndValidate()
}

func (c *Config) EncryptCredentials() error {
	if err := c.parseAndValidate(); err != nil {
		return err
	}

	credentials, ok := c.ProviderConfig.Credentials.(*Credentials)
	if !ok {
		return ErrInvalidCredentials
	}

	if err := credentials.Encrypt(c.crypto); err != nil {
		return err
	}

	c.ProviderConfig.Credentials = credentials
	return nil
}

func (c *Config) parseAndValidate() error {
	if c.valid {
		return nil
	}

	validationErrors := []error{}

	credentials, err := c.validateCredentials(c.ProviderConfig.Credentials)
	if err != nil {
		return err
	} else {
		c.ProviderConfig.Credentials = credentials
	}

	if len(c.ProviderConfig.Resources) != 1 {
		return ErrShouldHaveOneResource
	}
	r := c.ProviderConfig.Resources[0]
	if err := c.validateResourceConfig(r); err != nil {
		validationErrors = append(validationErrors, err)
	}

	if len(validationErrors) > 0 {
		errorStrings := []string{}
		for _, err := range validationErrors {
			errorStrings = append(errorStrings, err.Error())
		}
		return errors.New(strings.Join(errorStrings, "\n"))
	}

	c.valid = true
	return nil
}

func (c *Config) validateCredentials(value interface{}) (*Credentials, error) {
	var credentials Credentials
	if err := mapstructure.Decode(value, &credentials); err != nil {
		return nil, err
	}

	if err := c.validator.Struct(credentials); err != nil {
		return nil, err
	}

	saKeyJson, err := base64.StdEncoding.DecodeString(credentials.ServiceAccountKey)
	if err != nil {
		return nil, err
	}

	credentials.ServiceAccountKey = string(saKeyJson)

	return &credentials, nil
}

func (c *Config) validateResourceConfig(resource *domain.ResourceConfig) error {
	resourceTypeValidation := fmt.Sprintf("oneof=%s %s", ResourceTypeProject, ResourceTypeOrganization)
	if err := c.validator.Var(resource.Type, resourceTypeValidation); err != nil {
		return err
	}

	if len(resource.Roles) == 0 {
		return ErrRolesShouldNotBeEmpty
	}

	return nil
}

func (c *Config) validatePermissions(resource *domain.ResourceConfig, client GcloudIamClient) error {
	iamRoles, err := client.GetRoles()
	if err != nil {
		return err
	}

	var roles []*domain.Role
	for _, r := range iamRoles {
		roles = append(roles, &domain.Role{
			ID:          r.Name,
			Name:        r.Title,
			Description: r.Description,
		})
	}

	rolesMap := make(map[string]*domain.Role)
	for _, role := range roles {
		rolesMap[role.ID] = role
	}

	for _, ro := range resource.Roles {
		for _, p := range ro.Permissions {
			permission := fmt.Sprint(p)
			if _, ok := rolesMap[permission]; !ok {
				return fmt.Errorf("%w: %v", ErrInvalidProjectRole, permission)
			}
		}
	}

	return nil
}

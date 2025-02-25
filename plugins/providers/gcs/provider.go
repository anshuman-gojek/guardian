package gcs

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"cloud.google.com/go/iam"
	"github.com/mitchellh/mapstructure"
	"github.com/odpf/guardian/core/provider"
	"github.com/odpf/guardian/domain"
	"github.com/odpf/guardian/utils"
)

//go:generate mockery --name=Crypto --exported --with-expecter
type Crypto interface {
	domain.Crypto
}
type Provider struct {
	provider.UnimplementedClient
	provider.PermissionManager

	typeName string
	Clients  map[string]GCSClient
	crypto   Crypto
}

func NewProvider(typeName string, crypto Crypto) *Provider {
	return &Provider{
		typeName: typeName,
		Clients:  map[string]GCSClient{},
		crypto:   crypto,
	}
}

func (p *Provider) CreateConfig(pc *domain.ProviderConfig) error {
	c := NewConfig(pc, p.crypto)

	if err := c.parseAndValidate(); err != nil {
		return err
	}

	credentials, ok := c.ProviderConfig.Credentials.(*Credentials)
	if !ok {
		return ErrInvalidCredentialsType
	}

	if err := credentials.Encrypt(c.crypto); err != nil {
		return fmt.Errorf("encrypting creds: %w", err)
	}

	c.ProviderConfig.Credentials = credentials
	return nil
}

func (p *Provider) GetResources(pc *domain.ProviderConfig) ([]*domain.Resource, error) {
	var creds Credentials
	if err := mapstructure.Decode(pc.Credentials, &creds); err != nil {
		return nil, err
	}

	if err := creds.Decrypt(p.crypto); err != nil {
		return nil, err
	}

	client, err := p.getGCSClient(creds)
	if err != nil {
		return nil, err
	}

	resourceTypes := []string{}
	for _, rc := range pc.Resources {
		resourceTypes = append(resourceTypes, rc.Type)
	}

	var resources []*domain.Resource
	projectID := strings.Replace(creds.ResourceName, "projects/", "", 1)
	buckets, err := client.GetBuckets(context.TODO(), projectID)
	if err != nil {
		return nil, err
	}
	for _, b := range buckets {
		bucketResource := b.toDomain()
		bucketResource.ProviderType = pc.Type
		bucketResource.ProviderURN = pc.URN
		if utils.ContainsString(resourceTypes, ResourceTypeBucket) {
			resources = append(resources, bucketResource)
		}
	}

	return resources, nil
}

func (p *Provider) GetType() string {
	return p.typeName
}

func (p *Provider) GrantAccess(pc *domain.ProviderConfig, a domain.Grant) error {
	if err := validateProviderConfigAndAppealParams(pc, a); err != nil {
		return fmt.Errorf("invalid provider/appeal config: %w", err)
	}

	permissions := getPermissions(a)

	var creds Credentials
	if err := mapstructure.Decode(pc.Credentials, &creds); err != nil {
		return fmt.Errorf("error in decoding credentials%w", err)
	}

	if err := creds.Decrypt(p.crypto); err != nil {
		return fmt.Errorf("error in decrypting credentials%w", err)
	}

	client, err := p.getGCSClient(creds)
	if err != nil {
		return fmt.Errorf("error in getting new client: %w", err)
	}
	// identity is AccountType : AccountID, eg: "serviceAccount:test@email.com"
	identity := fmt.Sprintf("%s:%s", a.AccountType, a.AccountID)
	if a.Resource.Type == ResourceTypeBucket {
		b := new(Bucket)
		if err := b.fromDomain(a.Resource); err != nil {
			return fmt.Errorf("from Domain func error: %w", err)
		}
		for _, p := range permissions {
			role := iam.RoleName(string(p))
			if err := client.GrantBucketAccess(context.TODO(), *b, identity, role); err != nil {
				if errors.Is(err, ErrPermissionAlreadyExists) {
					return nil
				}
				return err
			}
		}
		return nil
	}
	return ErrInvalidResourceType
}

func (p *Provider) RevokeAccess(pc *domain.ProviderConfig, a domain.Grant) error {
	if err := validateProviderConfigAndAppealParams(pc, a); err != nil {
		return fmt.Errorf("invalid provider/appeal config: %w", err)
	}

	permissions := getPermissions(a)

	var creds Credentials
	if err := mapstructure.Decode(pc.Credentials, &creds); err != nil {
		return fmt.Errorf("error in decoding credentials%w", err)
	}

	if err := creds.Decrypt(p.crypto); err != nil {
		return fmt.Errorf("error in decrypting credentials%w", err)
	}

	client, err := p.getGCSClient(creds)
	if err != nil {
		return fmt.Errorf("error in getting new client: %w", err)
	}

	// identity is AccountType : AccountID, eg: "serviceAccount:test@email.com"
	identity := fmt.Sprintf("%s:%s", a.AccountType, a.AccountID)
	if a.Resource.Type == ResourceTypeBucket {
		b := new(Bucket)
		if err := b.fromDomain(a.Resource); err != nil {
			return fmt.Errorf("from Domain func error: %w", err)
		}
		for _, p := range permissions {
			var role iam.RoleName = iam.RoleName(string(p))
			if err := client.RevokeBucketAccess(context.TODO(), *b, identity, role); err != nil {
				if errors.Is(err, ErrPermissionAlreadyExists) {
					return nil
				}
				return err
			}
		}
		return nil
	}
	return ErrInvalidResourceType
}

func (p *Provider) GetRoles(pc *domain.ProviderConfig, resourceType string) ([]*domain.Role, error) {
	return provider.GetRoles(pc, resourceType)
}

func (p *Provider) GetAccountTypes() []string {
	return []string{AccountTypeUser, AccountTypeServiceAccount, AccountTypeGroup, AccountTypeDomain}
}

func validateProviderConfigAndAppealParams(pc *domain.ProviderConfig, a domain.Grant) error {
	if pc == nil {
		return ErrNilProviderConfig
	}
	if a.Resource == nil {
		return ErrNilResource
	}
	if a.Resource.ProviderType != pc.Type {
		return ErrProviderTypeMismatch
	}
	if a.Resource.ProviderURN != pc.URN {
		return ErrProviderURNMismatch
	}
	return nil
}

func getPermissions(a domain.Grant) []Permission {
	var permissions []Permission
	for _, p := range a.Permissions {
		permissions = append(permissions, Permission(p))
	}
	return permissions
}

func (p *Provider) getGCSClient(creds Credentials) (GCSClient, error) {
	projectID := strings.Replace(creds.ResourceName, "projects/", "", 1)
	if p.Clients[projectID] != nil {
		return p.Clients[projectID], nil
	}

	client, err := newGCSClient(projectID, []byte(creds.ServiceAccountKey))
	if err != nil {
		return nil, err
	}

	p.Clients[projectID] = client
	return client, nil
}

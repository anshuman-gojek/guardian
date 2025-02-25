package server

import (
	"errors"
	"fmt"

	"github.com/odpf/guardian/internal/store"
	"github.com/odpf/guardian/pkg/tracing"
	"github.com/odpf/guardian/plugins/notifiers"
	"github.com/odpf/salt/config"
)

type JobType string

const (
	FetchResources            JobType = "fetch_resources"
	ExpiringGrantNotification JobType = "grant_expiration_reminder"
	RevokeExpiredGrants       JobType = "grant_expiration_revocation"
)

type JobConfig struct {
	JobType JobType
	// Enabled is set as true for backward compatibility. If the job needs to be disabled, it must be present in the config with this value as false.
	Enabled  bool   `mapstructure:"enabled" default:"true"`
	Interval string `mapstructure:"interval"`
}

type Jobs struct {
	FetchResources             JobConfig `mapstructure:"fetch_resources"`
	RevokeExpiredGrants        JobConfig `mapstructure:"revoke_expired_grants"`
	ExpiringGrantNotification  JobConfig `mapstructure:"expiring_grant_notification"`
	RevokeExpiredAccess        JobConfig `mapstructure:"revoke_expired_access"`
	ExpiringAccessNotification JobConfig `mapstructure:"expiring_access_notification"`
}

type Config struct {
	Port                       int              `mapstructure:"port" default:"8080"`
	EncryptionSecretKeyKey     string           `mapstructure:"encryption_secret_key"`
	Notifier                   notifiers.Config `mapstructure:"notifier"`
	LogLevel                   string           `mapstructure:"log_level" default:"info"`
	DB                         store.Config     `mapstructure:"db"`
	AuthenticatedUserHeaderKey string           `mapstructure:"authenticated_user_header_key"`
	AuditLogTraceIDHeaderKey   string           `mapstructure:"audit_log_trace_id_header_key" default:"X-Trace-Id"`
	Jobs                       Jobs             `mapstructure:"jobs"`
	Telemetry                  tracing.Config   `mapstructure:"telemetry"`
}

func LoadConfig(configFile string) (Config, error) {
	var cfg Config
	loader := config.NewLoader(config.WithFile(configFile))

	if err := loader.Load(&cfg); err != nil {
		if errors.As(err, &config.ConfigFileNotFoundError{}) {
			fmt.Println(err)
			return cfg, nil
		}
		return Config{}, err
	}
	return cfg, nil
}

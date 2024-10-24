package configuration

import (
	"time"
)

type Configurations struct {
	App             AppProfile
	Service         ServicesProfile
	Database        DatabasesProfile
	DatabaseQueries DatabaseQueriesProfile
}

type (
	AppProfile struct {
		Config        RefreshConfig   `json:"config"`
		FeatureFlags  map[string]bool `json:"feature_flags"`
		MySQLDatabase DBConnection    `json:"mysql_database"`
	}
	DBConnection struct {
		DisableForeignKeyConstraintWhenMigrating bool          `json:"disable_foreign_key_constraint_when_migrating"`
		ConnectionName                           string        `json:"connection_name"`
		MaxRetries                               int           `json:"max_retries"`
		RetryIntervalTime                        time.Duration `json:"retry_interval_time"`
		LogQueries                               bool          `json:"log_queries"`
	}
)

type ServicesProfile struct {
	HTTPRestPool
}

type DatabasesProfile struct {
	MySQL
}

type DatabaseQueriesProfile struct {
	Queries
}

type BaseService struct {
	ServiceID string `json:"service_id"`
}

func (c *Configurations) GetFeatureFlag(flag string, fallback bool) bool {
	value, exists := c.App.FeatureFlags[flag]
	if !exists {
		return fallback
	}
	return value
}

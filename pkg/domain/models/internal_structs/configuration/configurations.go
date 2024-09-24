package configuration

import (
	"encoding/json"
	"fmt"
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
		SiteID                         string          `json:"site_id"`
		SitesEnabled                   string          `json:"sites_enabled"`
		Config                         RefreshConfig   `json:"config"`
		PipelinePreOptMaxIterations    int             `json:"pipeline_pre_opt_max_iterations"`
		FeatureFlags                   map[string]bool `json:"feature_flags"`
		MySQLRCBuffer                  DBConnection    `json:"mysql_rc_buffer"`
		DistributionOrderByFC          string          `json:"distribution_order_by_fc"`
		RehydrationFeatureFlagsGrouped map[string]bool `json:"rehydration_feature_flags_grouped"`
		NodesRCFC                      []string        `json:"nodes_rc_fc"`
	}
	DBConnection struct {
		ConnectionName    string        `json:"connection_name"`
		MaxRetries        int           `json:"max_retries"`
		RetryIntervalTime time.Duration `json:"retry_interval_time"`
		LogQueries        bool          `json:"log_queries"`
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

func (c *Configurations) GetFeatureFlagGrouped(flag string, fallback bool) bool {
	value, exists := c.App.RehydrationFeatureFlagsGrouped[flag]
	if !exists {
		return fallback
	}
	return value
}

func (c *Configurations) GetOrderFC() []string {
	var distributionMap map[string]int
	if err := json.Unmarshal([]byte(c.App.DistributionOrderByFC), &distributionMap); err != nil {
		fmt.Println("Error decoding JSON string:", err)
		return nil
	}

	orderFc := make([]string, len(distributionMap))
	for key, value := range distributionMap {
		orderFc[value-1] = key
	}
	return orderFc
}

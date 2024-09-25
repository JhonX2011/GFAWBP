package doubles

import (
	_ "embed"
	"encoding/json"

	mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"
	mic "github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/configuration"
)

const (
	QueryGetSomeInfo string = "SELECT * FROM `information`"
)

//go:embed test_data/config_profiles/app.json
var appConfigProfile string

//go:embed test_data/config_profiles/services.json
var servicesConfigProfile string

func GetAppConfigProfileInByte() []byte {
	return []byte(appConfigProfile)
}

func GetServicesConfigProfileInByte() []byte {
	return []byte(servicesConfigProfile)
}

func GetAppConfigProfileInMap() map[string]interface{} {
	var cfg map[string]interface{}
	if err := json.Unmarshal(GetAppConfigProfileInByte(), &cfg); err != nil {
		panic("GetAppConfigProfileInMap: " + err.Error())
	}

	return cfg
}

func GetRestClient() mic.RestClientConfig {
	var cfg mic.ServicesProfile
	if err := json.Unmarshal(GetServicesConfigProfileInByte(), &cfg); err != nil {
		panic("GetRestClient: " + err.Error())
	}

	if len(cfg.HTTPRestPool.Services) == 0 {
		return mic.RestClientConfig{}
	}

	return cfg.HTTPRestPool.Services[0]
}

func GetConfiguration() mic.Configurations {
	var cfg mic.Configurations
	_ = json.Unmarshal(GetAppConfigProfileInByte(), &cfg.App)
	_ = json.Unmarshal(GetServicesConfigProfileInByte(), &cfg.Service)

	return cfg
}

func GetResponseConfigurations() mcs.Configurations {
	return mcs.Configurations{
		Configs: []mcs.ConfigMember{
			{
				Name:  "Config1",
				Value: "Value1",
			},
			{
				Name:  "Config2",
				Value: "Value2",
			},
		},
	}
}

func GetDatabaseConfig(queryID string) *mic.Configurations {
	return &mic.Configurations{
		App: mic.AppProfile{
			MySQLDatabase: mic.DBConnection{
				MaxRetries:        3,
				RetryIntervalTime: 10,
			},
		},
		DatabaseQueries: mic.DatabaseQueriesProfile{
			Queries: mic.Queries{
				Queries: []mic.QueriesConfig{
					{
						QueryID:    queryID,
						QueryValue: QueryGetSomeInfo,
					},
				},
			},
		},
	}
}

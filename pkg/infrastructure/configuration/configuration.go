package configuration

import (
	"encoding/json"
	"sync"

	mic "github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/configuration"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/utils/config"
)

//nolint:gochecknoglobals
var (
	appProfile             = mic.AppProfileName.String()
	servicesProfile        = mic.ServicesProfileName.String()
	databasesProfile       = mic.DatabasesProfileName.String()
	databaseQueriesProfile = mic.DatabaseQueriesProfileName.String()
	unmarshal              = json.Unmarshal
)

type configuration struct {
	mutex  sync.RWMutex
	config *mic.Configurations
}

type Configuration interface {
	GetConfig() *mic.Configurations
	LoadConfig() error
	LoadJSONProfile(profileName string, mappingType interface{}) (interface{}, error)
}

func NewConfiguration() Configuration {
	return &configuration{}
}

func (c *configuration) LoadConfig() error {
	var config mic.Configurations

	if _, err := c.LoadJSONProfile(appProfile, &config.App); err != nil {
		return err
	}

	if _, err := c.LoadJSONProfile(servicesProfile, &config.Service); err != nil {
		return err
	}

	if _, err := c.LoadJSONProfile(databasesProfile, &config.Database); err != nil {
		return err
	}

	if _, err := c.LoadJSONProfile(databaseQueriesProfile, &config.DatabaseQueries); err != nil {
		return err
	}

	c.setConfig(&config)
	return nil
}

func (c *configuration) LoadJSONProfile(profileName string, mappingType interface{}) (interface{}, error) {
	bytes, err := readProfile(profileName)
	if err != nil {
		return nil, err
	}

	err = unmarshal(bytes, &mappingType)
	if err != nil {
		return nil, err
	}

	return &mappingType, nil
}

func (c *configuration) GetConfig() *mic.Configurations {
	if c.config == nil {
		return &mic.Configurations{}
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.config
}

func (c *configuration) setConfig(newConfig *mic.Configurations) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.config = newConfig
}

func readProfile(profileName string) ([]byte, error) {
	return config.Read(profileName)
}

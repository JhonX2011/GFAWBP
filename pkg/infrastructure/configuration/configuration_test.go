package configuration

import (
	"errors"
	"testing"

	mic "github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/configuration"
	gt "github.com/JhonX2011/GFAWBP/pkg/test/generic"
	"github.com/stretchr/testify/assert"
)

const configDir = "../../test/doubles/testdata/"

type configurationScenery struct {
	config Configuration
	gt.GenericTest
}

func givenConfigurationScenery() *configurationScenery {
	return &configurationScenery{
		config: NewConfiguration(),
	}
}

func (c *configurationScenery) givenEnvironmentVariables(t *testing.T) {
	t.Setenv("CONFIG_DIR", configDir)
}

func (c *configurationScenery) whenLoadConfigIsCall() {
	c.AError = c.config.LoadConfig()
}

func (c *configurationScenery) whenLoadJSONProfileIsCall(profileName string, mappingType interface{}) {
	c.AResult, c.AError = c.config.LoadJSONProfile(profileName, mappingType)
}

func (c *configurationScenery) whenGetConfigIsCall() {
	c.AResult = c.config.GetConfig()
}

func (c *configurationScenery) thenHaveAError(t *testing.T) {
	assert.NotNil(t, c.AError)
}

func TestGetConfigOk(t *testing.T) {
	s := givenConfigurationScenery()
	s.givenEnvironmentVariables(t)
	s.whenLoadConfigIsCall()
	s.whenGetConfigIsCall()
	s.ThenNotEmpty(t)
}

func TestGetConfigNull(t *testing.T) {
	s := givenConfigurationScenery()
	s.whenGetConfigIsCall()
	s.ThenNotEmpty(t)
	s.ThenEqual(t, &mic.Configurations{})
}

func TestLoadConfigOk(t *testing.T) {
	s := givenConfigurationScenery()
	s.givenEnvironmentVariables(t)
	s.whenLoadConfigIsCall()
	s.ThenNoHaveError(t)
}

func TestLoadConfigErrorAppProfile(t *testing.T) {
	oldValue := appProfile
	appProfile = "invalid_app_profile_name"
	defer func() { appProfile = oldValue }()
	s := givenConfigurationScenery()
	s.givenEnvironmentVariables(t)
	s.whenLoadConfigIsCall()
	s.thenHaveAError(t)
}

func TestLoadConfigErrorServicesProfile(t *testing.T) {
	oldValue := servicesProfile
	servicesProfile = "invalid_services_profile_name"
	defer func() { servicesProfile = oldValue }()
	s := givenConfigurationScenery()
	s.givenEnvironmentVariables(t)
	s.whenLoadConfigIsCall()
	s.thenHaveAError(t)
}

func TestLoadConfigErrorDatabasesProfile(t *testing.T) {
	oldValue := databasesProfile
	databasesProfile = "invalid_databases_profile_name"
	defer func() { databasesProfile = oldValue }()
	s := givenConfigurationScenery()
	s.givenEnvironmentVariables(t)
	s.whenLoadConfigIsCall()
	s.thenHaveAError(t)
}

func TestLoadConfigErrorDatabaseQueriesProfile(t *testing.T) {
	oldValue := databaseQueriesProfile
	databaseQueriesProfile = "invalid_databases_profile_name"
	defer func() { databaseQueriesProfile = oldValue }()
	s := givenConfigurationScenery()
	s.givenEnvironmentVariables(t)
	s.whenLoadConfigIsCall()
	s.thenHaveAError(t)
}

func TestLoadJSONProfileErrorUnmarshal(t *testing.T) {
	oldValue := unmarshal
	unmarshal = func(data []byte, v interface{}) error {
		return errors.New("json unmarshal error")
	}
	defer func() { unmarshal = oldValue }()
	s := givenConfigurationScenery()
	s.givenEnvironmentVariables(t)
	s.whenLoadJSONProfileIsCall(databasesProfile, nil)
	s.ThenErrorWithMessage(t, "json unmarshal error")
}

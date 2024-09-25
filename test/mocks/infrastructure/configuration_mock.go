package infrastructuremock

import (
	mic "github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/configuration"
	"github.com/stretchr/testify/mock"
)

type ConfigurationMock struct {
	mock.Mock
}

func (m *ConfigurationMock) GetConfig() *mic.Configurations {
	args := m.Called()
	return args.Get(0).(*mic.Configurations)
}

func (m *ConfigurationMock) LoadConfig() error {
	args := m.Called()
	return args.Error(0)
}

func (m *ConfigurationMock) LoadJSONProfile(profileName string, mappingType interface{}) (interface{}, error) {
	args := m.Called(profileName, mappingType)
	return args.Get(0), args.Error(1)
}

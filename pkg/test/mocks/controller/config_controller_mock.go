package controllermock

import (
	"github.com/stretchr/testify/mock"
)

type ConfigControllerMock struct {
	mock.Mock
}

func (m *ConfigControllerMock) RefreshConfiguration() error {
	ret := m.Called()

	if ret[0] != nil {
		return ret.Error(0)
	}
	return nil
}

func (m *ConfigControllerMock) GetConfigs() (interface{}, error) {
	ret := m.Called()

	if ret[1] != nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0), nil
}

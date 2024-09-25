package interactormock

import (
	"github.com/stretchr/testify/mock"
)

type ConfigInteractorMock struct {
	mock.Mock
}

func (_m *ConfigInteractorMock) Reload(_ int) error {
	ret := _m.Called()

	if ret[0] != nil {
		return ret.Error(0)
	}
	return nil
}

func (_m *ConfigInteractorMock) GetConfigurations() (interface{}, error) {
	ret := _m.Called()

	if ret[1] != nil {
		return nil, ret.Error(1)
	}
	return ret.Get(0), nil
}

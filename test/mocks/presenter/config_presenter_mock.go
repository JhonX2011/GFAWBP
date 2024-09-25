package presentermock

import (
	mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"
	"github.com/stretchr/testify/mock"
)

type ConfigPresenterMock struct {
	mock.Mock
}

func (c *ConfigPresenterMock) ResponseGetConfigs(_ []mcs.ConfigMember) mcs.Configurations {
	args := c.Called()
	return args[0].(mcs.Configurations)
}

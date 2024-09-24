package presentermock

import (
	mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"
	"github.com/stretchr/testify/mock"
)

type GetStructErrorPresenterMock struct {
	mock.Mock
}

func (c *GetStructErrorPresenterMock) LoadStructError(string, string, int, error, []mcs.ObjectErrors, bool) *mcs.MessageErrorResponse {
	args := c.Called()
	return args[0].(*mcs.MessageErrorResponse)
}

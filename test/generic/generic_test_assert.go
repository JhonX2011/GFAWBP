package generictest

import (
	"context"
	"net/http"
	"testing"

	stringu "github.com/JhonX2011/GFAWBP/pkg/infrastructure/utils/string"
	"github.com/stretchr/testify/assert"
)

type GenericTest struct {
	AError   error
	AResult  any
	RContext context.Context
}

func (g *GenericTest) GivenContext() {
	g.RContext = context.WithValue(context.Background(), stringu.RequestKey, "123")
}

func (g *GenericTest) ThenNotEmpty(t *testing.T) {
	assert.NotNil(t, g.AResult)
}

func (g *GenericTest) ThenEqual(t *testing.T, result any) {
	assert.Equal(t, result, g.AResult)
}

func (g *GenericTest) ThenErrorWithMessage(t *testing.T, msg string) {
	assert.NotNil(t, g.AError)
	assert.EqualError(t, g.AError, msg)
}

func (g *GenericTest) ThenErrorWithStatusCode(t *testing.T, expectedStatusCode int) {
	assert.NotNil(t, g.AError)
	assert.NotNil(t, g.AResult)
	assert.Equal(t, expectedStatusCode, g.AResult.(*http.Response).StatusCode)
}

func (g *GenericTest) ThenHaveError(t *testing.T) {
	assert.NotNil(t, g.AError)
}

func (g *GenericTest) ThenNoHaveError(t *testing.T) {
	assert.Nil(t, g.AError)
}

func (g *GenericTest) ThenEqualAny(t *testing.T, result any, expected any) {
	assert.Equal(t, expected, result)
}

func (g *GenericTest) ThenIsType(t *testing.T, expectedType, asserType any) {
	assert.IsType(t, expectedType, asserType)
}

package router

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	ut "github.com/JhonX2011/GFAWBP/pkg/infrastructure/utils/test"
	"github.com/JhonX2011/GFAWBP/test/mocks/controller"
	"github.com/JhonX2011/GOWebApplication/api"
	"github.com/stretchr/testify/assert"
)

type configRouteScenery struct {
	app                  *api.Application
	server               *httptest.Server
	configControllerMock *controllermock.ConfigControllerMock
	aError               error
	aResult              interface{}
}

func givenConfigRouteScenery() *configRouteScenery {
	configControllerMock := controllermock.ConfigControllerMock{}

	return &configRouteScenery{
		configControllerMock: &configControllerMock,
	}
}

func (c *configRouteScenery) givenConfigRouteApp() {
	app, err := api.NewWebApplication()
	if err != nil {
		panic("Error to create mock server")
	}
	r := NewConfigRoute(c.configControllerMock)
	rootRoutes := NewRouterRoot(app)
	rootRoutes.AddRoute(r)

	c.app = app
}

func (c *configRouteScenery) givenConfigRouteServer() {
	c.server = httptest.NewServer(c.app.Router)
}

func (c *configRouteScenery) whenRefreshConfigRequest(body io.Reader, err error) {
	url := fmt.Sprintf("%s/refresh_configs", c.server.URL)
	c.configControllerMock.On("RefreshConfiguration").Return(err)
	aResult, aError := ut.ExecuteRequest(url, body, http.MethodPost)
	c.aResult, c.aError = aResult, aError
	defer aResult.Body.Close() //nolint:gocritic
}

func (c *configRouteScenery) whenGetConfigRequest(err error) {
	url := fmt.Sprintf("%s/app_configs", c.server.URL)
	c.configControllerMock.On("GetConfigs").Return(nil, err)
	aResult, aError := ut.ExecuteRequest(url, nil, http.MethodGet)
	c.aResult, c.aError = aResult, aError
	defer aResult.Body.Close() //nolint:gocritic
}

func (c *configRouteScenery) thenNoError(t *testing.T, statusCode int) {
	assert.Nil(t, c.aError)
	assert.NotNil(t, c.aResult)
	assert.Equal(t, statusCode, c.aResult.(*http.Response).StatusCode)
}

func (c *configRouteScenery) thenError(t *testing.T) {
	assert.Nil(t, c.aError)
	assert.NotNil(t, c.aResult)
	assert.Equal(t, http.StatusInternalServerError, c.aResult.(*http.Response).StatusCode)
}

func (c *configRouteScenery) thenClose() {
	c.server.Close()
	if c.aResult.(*http.Response) != nil {
		c.aResult.(*http.Response).Body.Close()
	}
}

func (c *configRouteScenery) assertExpectations(t *testing.T) {
	c.configControllerMock.AssertExpectations(t)
}

func TestRefreshConfigResponseOK(t *testing.T) {
	s := givenConfigRouteScenery()
	s.givenConfigRouteApp()
	s.givenConfigRouteServer()
	s.whenRefreshConfigRequest(nil, nil)
	s.thenNoError(t, http.StatusOK)
	s.thenClose()
	s.assertExpectations(t)
}

func TestRefreshConfigResponseErrorController(t *testing.T) {
	s := givenConfigRouteScenery()
	s.givenConfigRouteApp()
	s.givenConfigRouteServer()
	s.whenRefreshConfigRequest(nil, errors.New("unable to collect data"))
	s.thenError(t)
	s.thenClose()
	s.assertExpectations(t)
}

func TestGetConfigResponseOK(t *testing.T) {
	s := givenConfigRouteScenery()
	s.givenConfigRouteApp()
	s.givenConfigRouteServer()
	s.whenGetConfigRequest(nil)
	s.thenNoError(t, http.StatusOK)
	s.thenClose()
	s.assertExpectations(t)
}

func TestGetConfigResponseErrorController(t *testing.T) {
	s := givenConfigRouteScenery()
	s.givenConfigRouteApp()
	s.givenConfigRouteServer()
	s.whenGetConfigRequest(errors.New("unable to collect data"))
	s.thenError(t)
	s.thenClose()
	s.assertExpectations(t)
}

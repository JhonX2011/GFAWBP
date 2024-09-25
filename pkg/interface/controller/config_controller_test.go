package controller

import (
	"errors"
	"net/http"
	"testing"

	mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"
	"github.com/JhonX2011/GFAWBP/test/doubles"
	"github.com/JhonX2011/GFAWBP/test/mocks/interactor"
	"github.com/JhonX2011/GFAWBP/test/mocks/presenter"
	"github.com/JhonX2011/GOWebApplication/utils/logger"
	ul "github.com/JhonX2011/GOWebApplication/utils/logger"
	"github.com/stretchr/testify/assert"
)

type configurationControllerScenery struct {
	mockConfigInteractor *interactormock.ConfigInteractorMock
	configController     ConfigController
	presenter            *presentermock.GetStructErrorPresenterMock
	aResult              any
	anError              error
}

func givenConfigurationControllerScenery(t *testing.T) *configurationControllerScenery {
	t.Parallel()
	mockConfigInteractor := interactormock.ConfigInteractorMock{}
	presenterMock := &presentermock.GetStructErrorPresenterMock{}
	return &configurationControllerScenery{
		mockConfigInteractor: &mockConfigInteractor,
		presenter:            presenterMock,
		configController:     NewConfigController(&mockConfigInteractor, logger.NewLogger(ul.DefaultOSExit), presenterMock),
	}
}

func (c *configurationControllerScenery) whenRefreshConfiguration(err error) {
	c.mockConfigInteractor.On("Reload").Return(err)
	c.anError = c.configController.RefreshConfiguration()
}

func (c *configurationControllerScenery) whenMockPresenterErrorIs(msg string) {
	c.presenter.On("LoadStructError").Return(doubles.GetLoadStructError(msg, msg, http.StatusFailedDependency, errors.New("err"), false))
}

func (c *configurationControllerScenery) whenGetConfigs(response interface{}, err error) {
	c.mockConfigInteractor.On("GetConfigurations").Return(response, err)
	c.aResult, c.anError = c.configController.GetConfigs()
}

func (c *configurationControllerScenery) thenNoError(t *testing.T) {
	assert.Nil(t, c.anError)
}

func (c *configurationControllerScenery) thenHaveAError(t *testing.T, msg string) {
	assert.NotNil(t, c.anError)
	assert.EqualError(t, c.anError, msg)
}

func (c *configurationControllerScenery) thenNoNil(t *testing.T, notNil any) {
	assert.NotNil(t, notNil)
}

func (c *configurationControllerScenery) thenEqual(t *testing.T, expectedMessage any, output any) {
	assert.Equal(t, expectedMessage, output)
}

func TestConfigControllerRefreshConfigurationOK(t *testing.T) {
	s := givenConfigurationControllerScenery(t)
	s.whenRefreshConfiguration(nil)
	s.thenNoError(t)
}

func TestConfigControllerRefreshConfigurationError(t *testing.T) {
	s := givenConfigurationControllerScenery(t)
	msg := "unable to refresh configuration"
	s.whenMockPresenterErrorIs(msg)
	s.whenRefreshConfiguration(errors.New(msg))
	s.thenHaveAError(t, msg)
}

func TestConfigControllerGetConfigurationsOK(t *testing.T) {
	expectedResult := mcs.ConfigMember{
		Name:  "someKey",
		Value: "someVal",
	}
	s := givenConfigurationControllerScenery(t)
	s.whenGetConfigs(expectedResult, nil)
	s.thenNoError(t)
	s.thenNoNil(t, s.aResult)

	converted, ok := s.aResult.(mcs.ConfigMember)
	assert.True(t, ok)
	s.thenEqual(t, expectedResult.Name, converted.Name)
	s.thenEqual(t, expectedResult.Value, converted.Value)
}

func TestConfigController_GetConfigurations_Error(t *testing.T) {
	s := givenConfigurationControllerScenery(t)
	msg := "unable to getConfigs"
	s.whenMockPresenterErrorIs(msg)
	s.whenGetConfigs(nil, errors.New(msg))
	s.thenNoNil(t, s.anError)
	s.thenHaveAError(t, msg)
}

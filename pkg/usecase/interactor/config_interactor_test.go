package interactor

import (
	"errors"
	"testing"

	mic "github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/configuration"
	"github.com/JhonX2011/GFAWBP/test/doubles"
	gt "github.com/JhonX2011/GFAWBP/test/generic"
	"github.com/JhonX2011/GFAWBP/test/mocks/infrastructure"
	"github.com/JhonX2011/GFAWBP/test/mocks/presenter"
	"github.com/stretchr/testify/mock"
)

const (
	initialRetry    int = 1
	defaultErrorMsg     = "any-error"
)

type configInteractorScenery struct {
	configInteractor    IConfigInteractor
	configMock          *infrastructuremock.ConfigurationMock
	configPresenterMock *presentermock.ConfigPresenterMock
	rBody               any
	gt.GenericTest
}

func givenConfigInteractorScenery() *configInteractorScenery {
	return &configInteractorScenery{
		configMock:          &infrastructuremock.ConfigurationMock{},
		configPresenterMock: &presentermock.ConfigPresenterMock{},
	}
}

func (c *configInteractorScenery) givenInteractor() {
	c.configInteractor = NewConfigInteractor(
		c.configMock,
		c.configPresenterMock,
	)
}

func (c *configInteractorScenery) givenBody(body any) {
	c.rBody = body
}

func (c *configInteractorScenery) whenMockGetConfigIsCall(cfg *mic.Configurations) {
	c.configMock.On("GetConfig").Return(cfg)
}

func (c *configInteractorScenery) whenMockLoadConfigIsCall(err error) {
	c.configMock.On("LoadConfig").Return(err)
}

func (c *configInteractorScenery) whenMockLoadJSONProfileIsCall(appConfig map[string]interface{}, err error) {
	c.configMock.On("LoadJSONProfile", mic.AppProfileName.String(), mock.Anything).Return(appConfig, err).
		Run(func(args mock.Arguments) {
			arg := args.Get(1).(*map[string]interface{})
			*arg = appConfig
		})
}

func (c *configInteractorScenery) whenMockResponseGetConfigsIsCall() {
	c.configPresenterMock.On("ResponseGetConfigs").Return(doubles.GetResponseConfigurations())
}

func (c *configInteractorScenery) whenReloadIsCall() {
	c.AError = c.configInteractor.Reload(c.rBody.(int))
}

func (c *configInteractorScenery) whenGetConfigurationsIsCall() {
	c.AResult, c.AError = c.configInteractor.GetConfigurations()
}

func TestReloadOk(t *testing.T) {
	cfg := doubles.GetConfiguration()
	s := givenConfigInteractorScenery()
	s.GivenContext()
	s.givenInteractor()
	s.givenBody(initialRetry)
	s.whenMockGetConfigIsCall(&cfg)
	s.whenMockLoadConfigIsCall(nil)
	s.whenReloadIsCall()
	s.ThenNoHaveError(t)
}

func TestReloadFail(t *testing.T) {
	cfg := mic.Configurations{}
	s := givenConfigInteractorScenery()
	s.GivenContext()
	s.givenInteractor()
	s.givenBody(5)
	s.whenMockGetConfigIsCall(&cfg)
	s.whenMockLoadConfigIsCall(errors.New(defaultErrorMsg))
	s.whenReloadIsCall()
	s.ThenErrorWithMessage(t, "unable to refresh configuration after [5] attempts")
}

func TestGetConfigurationsOk(t *testing.T) {
	s := givenConfigInteractorScenery()
	s.GivenContext()
	s.givenInteractor()
	s.whenMockLoadJSONProfileIsCall(doubles.GetAppConfigProfileInMap(), nil)
	s.whenMockResponseGetConfigsIsCall()
	s.whenGetConfigurationsIsCall()
	s.ThenNoHaveError(t)
}

func TestGetConfigurationsFail(t *testing.T) {
	var appConfig map[string]interface{}
	s := givenConfigInteractorScenery()
	s.GivenContext()
	s.givenInteractor()
	s.whenMockLoadJSONProfileIsCall(appConfig, nil)
	s.whenMockResponseGetConfigsIsCall()
	s.whenGetConfigurationsIsCall()
	s.ThenErrorWithMessage(t, "GetConfigurations: disabled")
}

func TestGetConfigurationsErrorLoad(t *testing.T) {
	var appConfig map[string]interface{}
	s := givenConfigInteractorScenery()
	s.GivenContext()
	s.givenInteractor()
	s.whenMockLoadJSONProfileIsCall(appConfig, errors.New(defaultErrorMsg))
	s.whenGetConfigurationsIsCall()
	s.ThenErrorWithMessage(t, defaultErrorMsg)
}

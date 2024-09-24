package controller

import (
	"net/http"

	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/logger"
	"github.com/JhonX2011/GFAWBP/pkg/interface/presenter"
	"github.com/JhonX2011/GFAWBP/pkg/usecase/interactor"
)

const initialRetry int = 1

type configController struct {
	configInteractor interactor.IConfigInteractor
	varLog           logger.Logger
	errorPresenter   presenter.IGetStructErrorPresenter
}

type ConfigController interface {
	RefreshConfiguration() error
	GetConfigs() (interface{}, error)
}

func NewConfigController(i interactor.IConfigInteractor, l logger.Logger, pe presenter.IGetStructErrorPresenter) ConfigController {
	return &configController{
		configInteractor: i,
		varLog:           l,
		errorPresenter:   pe,
	}
}

func (p *configController) RefreshConfiguration() error {
	p.varLog.Info("Refreshing configuration")
	if err := p.configInteractor.Reload(initialRetry); err != nil {
		p.varLog.Error("Refreshing config", err)
		return p.errorPresenter.LoadStructError("REFRESH_ERROR", err.Error(), http.StatusInternalServerError, err, nil,
			presenter.IsRetryableFalse)
	}
	p.varLog.Info("Configuration reloaded OK")
	return nil
}

func (p *configController) GetConfigs() (interface{}, error) {
	response, err := p.configInteractor.GetConfigurations()
	if err != nil {
		p.varLog.Error("GetConfigs", err)
		return nil, p.errorPresenter.LoadStructError("DEBUG_DISABLED", err.Error(), http.StatusInternalServerError, err,
			nil, presenter.IsRetryableFalse)
	}

	return response, nil
}

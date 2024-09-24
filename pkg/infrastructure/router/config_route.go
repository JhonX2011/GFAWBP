package router

import (
	"net/http"

	mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/api/web"
	"github.com/JhonX2011/GFAWBP/pkg/interface/controller"
)

type configRoute struct {
	configController controller.ConfigController
}

func NewConfigRoute(i controller.ConfigController) Route {
	return &configRoute{i}
}

func (p *configRoute) Config(group *web.RouteGroup) {
	group.Post("/refresh_config", p.refreshConfig)
	group.Get("/app_configs", p.getConfigsHandler)
}

func (p *configRoute) refreshConfig(w http.ResponseWriter, _ *http.Request) error {
	if err := p.configController.RefreshConfiguration(); err != nil {
		errorMessage := err.(*mcs.MessageErrorResponse)
		return web.NewError(errorMessage.Code, errorMessage.Message)
	}

	return web.EncodeJSON(w, "OK", http.StatusOK)
}

func (p *configRoute) getConfigsHandler(w http.ResponseWriter, _ *http.Request) error {
	response, err := p.configController.GetConfigs()

	if err != nil {
		errorMessage := err.(*mcs.MessageErrorResponse)
		return web.NewError(errorMessage.Code, errorMessage.Message)
	}

	return web.EncodeJSON(w, response, http.StatusOK)
}

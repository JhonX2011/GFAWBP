package registry

import (
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/router"
	ic "github.com/JhonX2011/GFAWBP/pkg/interface/controller"
	ip "github.com/JhonX2011/GFAWBP/pkg/interface/presenter"
	ui "github.com/JhonX2011/GFAWBP/pkg/usecase/interactor"
	up "github.com/JhonX2011/GFAWBP/pkg/usecase/presenter"
)

func (r *registry) NewConfigRoute() router.Route {
	return router.NewConfigRoute(r.NewConfigController())
}

func (r *registry) NewConfigController() ic.ConfigController {
	return ic.NewConfigController(r.NewConfigInteractor(), r.log, r.NewGetGetStructErrorPresenter())
}

func (r *registry) NewConfigInteractor() ui.IConfigInteractor {
	return ui.NewConfigInteractor(r.config, r.NewGetConfigsPresenter())
}

func (r *registry) NewGetConfigsPresenter() up.IGetConfigsPresenter {
	return ip.NewGetConfigsPresenter()
}

func (r *registry) NewGetGetStructErrorPresenter() ip.IGetStructErrorPresenter {
	return ip.NewGetGetStructErrorPresenter()
}

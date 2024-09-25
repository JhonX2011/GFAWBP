package router

import (
	"net/http"

	"github.com/JhonX2011/GOWebApplication/api"
	"github.com/JhonX2011/GOWebApplication/api/web"
)

type router struct {
	routes *web.RouteGroup
	app    *api.Application
}

type Route interface {
	Config(*web.RouteGroup)
}

type Router interface {
	AddRoute(Route)
}

func (r *router) AddRoute(route Route) {
	route.Config(r.routes)
}

func NewRouterRoot(app *api.Application) Router {
	routes := app.Group("/")
	return &router{app: app, routes: routes}
}

func NewRouter(app *api.Application) Router {
	routes := app.Group("/api/v1")
	routes.Get("/", notFoundHandler)
	routes.Post("/", notFoundHandler)

	return &router{app: app, routes: routes}
}

func notFoundHandler(w http.ResponseWriter, _ *http.Request) error {
	return web.EncodeJSON(w, "", http.StatusNotFound)
}

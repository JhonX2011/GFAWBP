package router

import (
	"net/http/httptest"
	"testing"

	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/api"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/api/web"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/logger"
	ul "github.com/JhonX2011/GFAWBP/pkg/infrastructure/utils/logger"
	"github.com/stretchr/testify/assert"
)

type routerScenery struct {
	routerRoot   Router
	routerClient Router
	routeTest    Route
	anError      error
}

func givenRouterScenery() *routerScenery {
	l := logger.NewLogger(ul.DefaultOSExit)
	app, _ := api.NewWebApplication(l)
	return &routerScenery{
		routerClient: NewRouter(app),
		routerRoot:   NewRouterRoot(app),
	}
}

func (r *routerScenery) givenRoute(route Route) {
	r.routeTest = route
}

func (r *routerScenery) whenAddRouteIsExecuted() {
	r.routerClient.AddRoute(r.routeTest)
}

func (r *routerScenery) whenAddRouteRootIsExecuted() {
	r.routerRoot.AddRoute(r.routeTest)
}

func (r *routerScenery) whenNotFoundIsExecuted() {
	r.anError = notFoundHandler(httptest.NewRecorder(), nil)
}

func (r *routerScenery) thenNoError(t *testing.T) {
	assert.Nil(t, r.anError)
}

type testRoute struct {
	handler web.Handler
}

func NewTestRoute(handler web.Handler) Route {
	return &testRoute{handler: handler}
}

func (p *testRoute) Config(group *web.RouteGroup) {
	group.Post("/test-route", p.handler)
}

func TestAddOK(t *testing.T) {
	s := givenRouterScenery()
	t.Parallel()
	s.givenRoute(NewTestRoute(nil))
	s.whenAddRouteIsExecuted()
}

func TestAddOKRoot(t *testing.T) {
	s := givenRouterScenery()
	t.Parallel()
	s.givenRoute(NewTestRoute(nil))
	s.whenAddRouteRootIsExecuted()
}

func TestAddOKWithHandler(t *testing.T) {
	s := givenRouterScenery()
	t.Parallel()
	s.givenRoute(NewTestRoute(notFoundHandler))
	s.whenAddRouteRootIsExecuted()
}

func TestNotFoundHandlerHandler(t *testing.T) {
	s := givenRouterScenery()
	t.Parallel()
	s.whenNotFoundIsExecuted()
	s.thenNoError(t)
}

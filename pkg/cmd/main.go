package main

import (
	"log"

	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/api"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/configuration"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/initializer"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/logger"
	appRegistry "github.com/JhonX2011/GFAWBP/pkg/infrastructure/registry"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/router"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/utils/environment"
	ul "github.com/JhonX2011/GFAWBP/pkg/infrastructure/utils/logger"
)

func init() { //nolint:gochecknoinits
	environment.LoadEnvironment()

	environment.PrintEnv()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	l := logger.NewLogger(ul.DefaultOSExit)

	app, err := api.NewWebApplication(l)
	if err != nil {
		l.Error("Unable to start web application", err)
		return err
	}

	registry, err := setupServices(l)
	if err != nil {
		return err
	}

	registerRoutes(app, registry)

	return app.Run()
}

func setupServices(l logger.Logger) (appRegistry.Registry, error) {
	configClient, err := initializer.InitConfigurationClient()
	if err != nil {
		return nil, err
	}

	return createRegistry(
		l,
		configClient,
	)
}

func createRegistry(
	l logger.Logger,
	configClient configuration.Configuration,
) (appRegistry.Registry, error) {

	return appRegistry.NewRegistry(
		l,
		configClient,
	), nil
}

func registerRoutes(app *api.Application, r appRegistry.Registry) {
	// -- Initialize routes
	rootRoutes := router.NewRouterRoot(app)

	// -- Root routes
	rootRoutes.AddRoute(r.NewConfigRoute())

	// -- Business routes
	//businessRoutes := router.NewRouter(app)
	//businessRoutes.AddRoute(r.NewGetDistributionRoute())
}

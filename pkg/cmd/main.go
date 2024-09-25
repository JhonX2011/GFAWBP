package main

import (
	"log"

	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/configuration"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/database/gorm"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/initializer"
	appRegistry "github.com/JhonX2011/GFAWBP/pkg/infrastructure/registry"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/router"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/utils/environment"
	"github.com/JhonX2011/GOWebApplication/api"
	"github.com/JhonX2011/GOWebApplication/utils/logger"
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
	app, err := api.NewWebApplication()
	if err != nil {
		app.Logger.Error("Unable to start web application", err)
		return err
	}

	registry, err := setupServices(app.Logger)
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

	gormClient, err := initializer.InitDatabaseMySQLClient(configClient)
	if err != nil {
		return nil, err
	}

	return createRegistry(
		l,
		configClient,
		gormClient,
	)
}

func createRegistry(
	l logger.Logger,
	configClient configuration.Configuration,
	gormClient gorm.IClientGorm,
) (appRegistry.Registry, error) {

	return appRegistry.NewRegistry(
		l,
		configClient,
		gormClient,
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

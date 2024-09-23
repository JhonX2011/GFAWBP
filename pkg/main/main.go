package main

import (
	"log"

	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/utils/environment"
)

func init() { //nolint:gochecknoinits
	// -- Load Environment Vars
	environment.LoadEnvironment()

	// -- Print Environment Vars
	environment.PrintEnv()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//l := logger.NewLogger(ul.DefaultOSExit)
	//app, err := fury.NewWebApplication(fury.WithEnableProfiling())
	//if err != nil {
	//	l.Error("Unable to start web application", err)
	//	return err
	//}
	//
	//registry, err := setupServices(app, l)
	//if err != nil {
	//	return err
	//}
	//
	//registerRoutes(app, registry)
	//
	//uHTTP.NotFoundHandler(app, umap.GetGenericMessageError("Route not found", http.StatusNotFound), l)
	//
	//return app.Run()

	return nil
}

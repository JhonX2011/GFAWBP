package api

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/api/web"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/logger"
)

const (
	_defaultWebApplicationPort = "8080"
	_defaultNetworkProtocol    = "tcp"
)

type Application struct {
	*web.Router

	logger  logger.Logger
	address string
}

func NewWebApplication(logger logger.Logger) (*Application, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = _defaultWebApplicationPort
	}

	address := ":" + port

	listener, err := net.Listen(_defaultNetworkProtocol, address)
	if err != nil {
		log.Fatalf("The provided port [%s] is not available: %v", address, err)
		return nil, err
	}
	logger.Info("Running application | address", address)
	defer listener.Close()

	return &Application{
		Router:  web.New(),
		logger:  logger,
		address: address,
	}, nil
}

func (a *Application) Run() error {
	a.defaultRoutes()

	srv := &http.Server{
		Addr:         a.address,
		Handler:      a.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (a *Application) defaultRoutes() {
	a.Router.Get("/ping", func(w http.ResponseWriter, r *http.Request) error {
		return web.EncodeJSON(w, "pong", 200)
	})
}

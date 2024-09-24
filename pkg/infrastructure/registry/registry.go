package registry

import (
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/configuration"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/logger"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/router"
)

type registry struct {
	log    logger.Logger
	config configuration.Configuration
}

type Registry interface {
	NewConfigRoute() router.Route
}

func NewRegistry(
	l logger.Logger,
	c configuration.Configuration,
) Registry {

	return &registry{
		log:    l,
		config: c,
	}
}

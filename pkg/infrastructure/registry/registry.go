package registry

import (
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/configuration"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/database/gorm"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/router"
	"github.com/JhonX2011/GOWebApplication/utils/logger"
)

type registry struct {
	log        logger.Logger
	config     configuration.Configuration
	gormClient gorm.IClientGorm
}

type Registry interface {
	NewConfigRoute() router.Route
}

func NewRegistry(
	l logger.Logger,
	c configuration.Configuration,
	gormClient gorm.IClientGorm,
) Registry {

	return &registry{
		log:        l,
		config:     c,
		gormClient: gormClient,
	}
}

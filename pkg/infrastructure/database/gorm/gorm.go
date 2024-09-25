package gorm

import (
	"context"
	"log"
	"os"
	"time"

	mic "github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/configuration"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type gormClient struct {
	db     *gorm.DB
	config *mic.DBConnection
}

type (
	IClientGorm interface {
		ITransaction
		GetDB(context.Context) *gorm.DB
		RetryQuery(context.Context, func() *gorm.DB) *gorm.DB
	}
	ITransaction interface {
		Begin(context.Context) (context.Context, func(), error)
		Commit(context.Context) error
		Rollback(context.Context) error
	}
)

func NewGormClient(dialector gorm.Dialector, config *mic.DBConnection) (IClientGorm, error) {
	gormDB, err := gorm.Open(dialector, getGormConfig(config))
	if err != nil {
		return nil, err
	}

	return &gormClient{
		db:     gormDB,
		config: config,
	}, nil
}

func (gc *gormClient) RetryQuery(ctx context.Context, queryFunc func() *gorm.DB) *gorm.DB {
	ticker := time.NewTicker(gc.config.RetryIntervalTime * time.Second)
	defer ticker.Stop()
	var result *gorm.DB

	for i := 0; i < gc.config.MaxRetries; i++ {
		result = queryFunc()
		if result.Error == nil {
			return result
		}

		if i == gc.config.MaxRetries-1 {
			break
		}
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return result
		}
	}

	return result
}

func getGormConfig(config *mic.DBConnection) *gorm.Config {
	logLevel := logger.Error
	if config.LogQueries {
		logLevel = logger.Info
	}

	return &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logLevel,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      false,
				Colorful:                  false,
			},
		),
	}
}

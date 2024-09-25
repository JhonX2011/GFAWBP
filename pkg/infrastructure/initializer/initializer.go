package initializer

import (
	"encoding/json"
	"errors"

	mic "github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/configuration"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/configuration"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/database/gorm"
	db "github.com/JhonX2011/GFAWBP/pkg/infrastructure/database/mysql"
	"gorm.io/driver/mysql"
	g "gorm.io/gorm"
)

func InitConfigurationClient() (configuration.Configuration, error) {
	cfg := configuration.NewConfiguration()
	if cfg == nil {
		return nil, errors.New("unable to load configuration")
	}

	// Loads all configuration profiles
	if err := cfg.LoadConfig(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func InitDatabaseMySQLClient(config configuration.Configuration) (gorm.IClientGorm, error) {
	cfg := config.GetConfig()
	service, err := cfg.Database.MySQL.Get(mic.MySqlDB)
	if err != nil {
		return nil, err
	}

	dialector, err := initMySQLDialector(&service, cfg.App.MySQLDatabase.ConnectionName)
	if err != nil {
		return nil, err
	}

	ormDB, err := initGormClient(dialector, &cfg.App.MySQLDatabase)
	if err != nil {
		return nil, err
	}

	return ormDB, nil
}

func initMySQLDialector(mySQLConfig *mic.MySQLConfig, connectionName string) (g.Dialector, error) {
	var conn db.IDBClient
	var dialector g.Dialector
	var err error

	bytes, _ := json.Marshal(mySQLConfig)
	conn, err = db.NewMysqlConn(bytes)
	if err != nil {
		return nil, err
	}

	sqlDB, err := conn.Get(connectionName)
	if err != nil {
		return nil, err
	}

	dialector = mysql.New(mysql.Config{
		Conn: sqlDB,
	})

	return dialector, err
}

func initGormClient(sqlDialector g.Dialector, config *mic.DBConnection) (gorm.IClientGorm, error) {
	return gorm.NewGormClient(sqlDialector, config)
}

//func InitRestClient(cfg configuration.Configuration, restClientName string, metricsCoreClient metricscore.IMetricsCoreClient) (httprest.IHttpRest, error) { //nolint:lll
//	service, err := cfg.GetConfig().Service.HTTPRestPool.Get(restClientName)
//	if err != nil {
//		return nil, err
//	}
//
//	return httprest.NewHTTPRestProducer(metricsCoreClient).Produce(&service)
//}
//
//func InitMySQLDialector(mySQLConfig *mic.MySQLConfig, connectionName string) (g.Dialector, error) {
//	var conn database.IDBClient
//	var dialector g.Dialector
//	var err error
//
//	bytes, _ := json.Marshal(mySQLConfig)
//	conn, err = database.NewMysqlConn(bytes)
//	if err != nil {
//		return nil, err
//	}
//
//	sqlDB, err := conn.Get(connectionName)
//	if err != nil {
//		return nil, err
//	}
//
//	dialector = mysql.New(mysql.Config{
//		Conn: sqlDB,
//	})
//
//	return dialector, err
//}
//
//func InitGormClient(sqlDialector g.Dialector, config *mic.DBConnection) (gorm.IClientGorm, error) {
//	return gorm.NewGormClient(sqlDialector, config)
//}

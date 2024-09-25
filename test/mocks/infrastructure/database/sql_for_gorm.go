package databasemock

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MockSQLForGormConnection struct {
}

func GetDB() (*sql.DB, sqlmock.Sqlmock, func(), error) {
	db, mockSQL, err := sqlmock.New()
	if err != nil {
		return nil, nil, nil, err
	}

	return db, mockSQL, func() {
		err := db.Close()
		if err != nil {
			return
		}
	}, nil
}

func GetDialector(db *sql.DB) gorm.Dialector {
	return mysql.New(mysql.Config{
		Conn:                      db,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})
}

func GetGormDB(db *sql.DB) (*gorm.DB, error) {
	dialector := GetDialector(db)
	return gorm.Open(dialector, &gorm.Config{})
}

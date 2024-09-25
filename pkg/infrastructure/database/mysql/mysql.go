package mysql

import (
	"database/sql"
	"encoding/json"

	"github.com/JhonX2011/GOWebApplication/database/mysqlconnect"
)

type mySQLConn struct {
	connections mysqlconnect.Connections
}

type IDBClient interface {
	Get(name string) (*sql.DB, error)
	List() []*sql.DB
	Close() error
}

func NewMysqlConn(configJSON []byte) (IDBClient, error) {
	var config mysqlconnect.Config
	if err := json.Unmarshal(configJSON, &config); err != nil {
		return nil, err
	}

	connections, err := mysqlconnect.Open(config)
	if err != nil {
		return nil, err
	}
	return &mySQLConn{connections: connections}, nil
}

func (m *mySQLConn) Get(name string) (*sql.DB, error) {
	return m.connections.Get(name)
}

func (m *mySQLConn) List() []*sql.DB {
	return m.connections.List()
}

func (m *mySQLConn) Close() error {
	return m.connections.Close()
}

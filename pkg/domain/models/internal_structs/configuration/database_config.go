package configuration

import "fmt"

type MySQLConfig struct {
	BaseService
	Dsn         string `json:"dsn,omitempty"`
	Cluster     string `json:"cluster,omitempty"`
	Schema      string `json:"schema,omitempty"`
	Connections []struct {
		Name           string `json:"name,omitempty"`
		IsMaster       bool   `json:"is_master,omitempty"`
		IsReadOnly     bool   `json:"is_read_only,omitempty"`
		Parameters     string `json:"parameters,omitempty"`
		ConnectionPool struct {
			ConnMaxLifetime    string `json:"conn_max_lifetime,omitempty"`
			MaxIdleConnections int    `json:"max_idle_connections,omitempty"`
			MaxOpenConnections int    `json:"max_open_connections,omitempty"`
			ConnMaxIdleTime    string `json:"conn_max_idle_time,omitempty"`
		} `json:"connection_pool"`
	} `json:"connections"`
}

type MySQL struct {
	Services []MySQLConfig `json:"mysql"`
}

func (s MySQL) Get(serviceID string) (MySQLConfig, error) {
	for _, service := range s.Services {
		if service.ServiceID == serviceID {
			return service, nil
		}
	}

	return MySQLConfig{}, fmt.Errorf("service not found [%s]", serviceID)
}

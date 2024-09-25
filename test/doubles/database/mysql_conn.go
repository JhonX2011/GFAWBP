package database

import _ "embed"

//go:embed test_data/wrong_mysql_conn_string.json
var mysqlConnString string

//go:embed test_data/mysql_conn_string.json
var mysqlConnStringOK string

func GetMysqlConnString() []byte {
	return []byte(mysqlConnString)
}

func GetMysqlConnStringOK() []byte {
	return []byte(mysqlConnStringOK)
}

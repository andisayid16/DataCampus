package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "NUSEB1jLGkRnCc1ZFYWVRpjzqSYqbVnh"
	dbName := "railway" // replace with your database name
	dbHost := "roundhouse.proxy.rlwy.net"
	dbPort := "13023"

	dbURL := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	db, err := sql.Open(dbDriver, dbURL)
	return db, err
}

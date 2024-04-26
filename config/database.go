package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "NwWQVNmDDOC1AFeFldKjgjhJHEF1vDxS"
	dbName := "DataCampus"
	dbHost := "viaduct.proxy.rlwy.net"
	dbPort := "37408"

	dbURL := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	db, err := sql.Open(dbDriver, dbURL)
	return db, err
}

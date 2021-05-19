package config

import (
	"database/sql"
	
	_ "github.com/go-sql-driver/mysql"
)

func GetMySQLDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbHost := "localhost"
	dbPort := "3306"
	dbUser	:= "root"
	dbPass := ""
	dbName := "golang"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	return 
}
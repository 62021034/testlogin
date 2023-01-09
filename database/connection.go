package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	driver   = "postgres"
	host     = "127.0.0.1"
	portData = 5432
	user     = "postgres"
	password = "111111"
	dbname   = "TestLoginEP2"
)

func Connect() error {
	var err error
	DB, err = sql.Open(driver, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, portData, user, password, dbname))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}
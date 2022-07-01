package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateDBConnection() (*sql.DB, error) {

	connectionString := "host=localhost port=5432 user=postgres password=root dbname=alingan sslmode=disable"

	con, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatalf("Error connect to DB %v ", err.Error())
	}

	return con, nil

}

package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func CreateDBConnection() (*sql.DB, error) {

	connectionString := "host=localhost port=5432 user=postgres password=root dbname=alingan sslmode=disable"

	con, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return con, nil

}

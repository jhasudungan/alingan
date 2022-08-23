package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateDBConnection() (*sql.DB, error) {

	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbHostPort := os.Getenv("DB_HOST_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		dbHost,
		dbHostPort,
		dbUser,
		dbPass,
		dbName)

	con, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return con, nil

}

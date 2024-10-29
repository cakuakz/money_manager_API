package config

import (
	"database/sql"
	"fmt"
	"os"

	"money-manager/constants"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetUpDatabaseConnection() *sql.DB {
	if os.Getenv("APP_ENV") != constants.ENUM_RUN_PRODUCTION {
		err := godotenv.Load("config/.env")
		if err != nil {
			panic(err)
		}
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`, dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}

func CloseDatabaseConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
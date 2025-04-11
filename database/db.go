package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var DB *sql.DB
var err error
var DBDriver = "postgres"

func Connect() {
	// Initialize the database connection
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Connect to the database
	DB, err = sql.Open(DBDriver, connectionString)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	// Check if the connection is successful

	if err = DB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}
}

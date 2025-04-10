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
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database unreachable:", err)
	}

	// Create the bioskop table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS bioskop (
		id SERIAL PRIMARY KEY,
		nama VARCHAR(255) NOT NULL,
		lokasi VARCHAR(255) NOT NULL,
		rating FLOAT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create bioskop table:", err)
	}
}

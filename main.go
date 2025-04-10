package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thomaspanji/go-simple-web-app-bioskop/database"
	"github.com/thomaspanji/go-simple-web-app-bioskop/handlers"
)

func main() {
	// Load the environment variables from the .env file
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize the database connection
	database.Connect()
	// Initialize the router
	router := gin.Default()
	// Set up routes
	router.POST("/bioskop", handlers.CreateBioskopHandler)
	// Start the server
	// Listen and serve on port 8080
	router.Run(":8080")

}

// @title           My Bus Service API
// @version         1.0
// @description     This API serves as an interface to interact with the My Bus Service platform, providing endpoints for managing bus routes, bookings, and user interactions.
// @termsOfService  http://swagger.io/terms/

// @contact.name    Mohammad Mizan
// @contact.url     http://swagger.io/support
// @contact.email   takbir.jcd@gmail.com

// @license.name    Apache License Version 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8081
// @BasePath        /api/v1
package main

import (
	"bus-service/internal/api"
	"bus-service/internal/config"
	"bus-service/internal/models"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading environment variables from system")
	}
	database := config.NewDatabase(&models.Bus{})
	defer database.Close()

	// Get the port number from the environment variable.
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT environment variable is required and not set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	// Create and start the server.
	server := api.NewServer(database)
	server.Start(":" + strconv.Itoa(port)) // Assuming Start will handle errors internally
}

package main

import (
	"github.com/joho/godotenv"
	"log"
	"notification-service/internal/api"
	"notification-service/internal/config"
	"notification-service/internal/models"
	"os"
	"strconv"
)

// @contact.name  MOHAMMAD IZANAMI RAHMAN
// @contact.email  takbir.jcd@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading environment variables from system")
	}
	database := config.NewDatabase(&models.Notification{}, &models.UserNotificationPreferences{})
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

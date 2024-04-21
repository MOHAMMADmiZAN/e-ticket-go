// @title           My Auth Service API
// @version         1.0
// @description     The My Auth Service API provides a secure and efficient means for handling authentication and authorization processes. By exposing endpoints for user registration, authentication, and account management, this API facilitates the establishment and maintenance of secure user sessions and the protection of user data. The API is meticulously designed to offer robust security features that align with industry best practices.

// @termsOfService  http://swagger.io/terms/

// @contact.name    Mohammad Mizan
// @contact.url     http://swagger.io/support
// @contact.email   takbir.jcd@gmail.com

// @license.name    Apache License Version 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8083
// @BasePath        /api/v1/auth

package main

import (
	"auth-service/internal/api"
	"auth-service/internal/config"
	"auth-service/internal/models"
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
	database := config.NewDatabase(&models.User{}, &models.LoginHistory{}, &models.UserVerification{})
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

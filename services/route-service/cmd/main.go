package main

import (
	"e-ticket/services/route-service/internal/api"
	"e-ticket/services/route-service/internal/config"
	"e-ticket/services/route-service/internal/model"
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
	database := config.NewDatabase(&model.Route{}, &model.Stop{}, &model.Schedule{})
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

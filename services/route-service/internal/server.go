package internal

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"route-service/internal/api/v1"
	"route-service/internal/config"
	"route-service/pkg/middleware"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	Router *gin.Engine
	DB     *config.Database
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(databaseClient *config.Database) *Server {
	// Set the router as the default gin router
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger(), middleware.ErrorHandlingMiddleware(), middleware.SwaggerHost()) // Add Logger middleware

	s := &Server{
		Router: r,
		DB:     databaseClient,
	}
	// Setup V1 routes
	v1.SetupRoutes(s.Router, s.DB)
	return s
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: s.Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server started successfully on", addr)

	// Wait for interrupt signal to gracefully shut down the server with a timeout.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Shortened timeout
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

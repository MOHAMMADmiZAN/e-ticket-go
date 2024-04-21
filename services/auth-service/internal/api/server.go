package api

import (
	_ "auth-service/docs" // Required for Swagger docs
	"auth-service/internal/api/handler"
	"auth-service/internal/api/middleware"
	"auth-service/internal/config"
	"auth-service/internal/repository"
	"auth-service/internal/services"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	Router *gin.Engine
	DB     *config.Database
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(databaseClient *config.Database) *Server {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger(), middleware.ErrorHandlingMiddleware()) // Add Logger middleware
	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))
	s := &Server{
		Router: r,
		DB:     databaseClient,
	}
	s.routes()
	return s
}

// routes registers all the routes to the router.
func (s *Server) routes() {

	// API Versioning
	v1 := s.Router.Group("/api/v1/auth")

	// Setup user handlers
	u := handler.NewUserHandler(services.NewUserService(repository.NewUserRepository(s.DB.Conn)))

	// Setup user routes
	s.setupUserRoutes(v1, u)

	// Health check route
	s.setupHealthCheckRoute()

	// Catch-all route for handling unmatched routes (404 Not Found)
	s.setupNoRouteHandler()
}

func (s *Server) setupHealthCheckRoute() {
	s.Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Auth services is running"})
	})
}

func (s *Server) setupNoRouteHandler() {
	s.Router.NoRoute(func(c *gin.Context) {
		// Improved error message
		c.JSON(http.StatusNotFound, gin.H{
			"error": gin.H{
				"code":    http.StatusNotFound,
				"type":    "PAGE_NOT_FOUND",
				"message": "The requested resource was not found on this server.",
			},
			"details": "Check the URL for errors or contact support if the problem persists.",
		})
	})
}

func (s *Server) setupUserRoutes(v1 *gin.RouterGroup, u *handler.UserHandler) {
	v1.POST("/register", u.RegisterUser)
	v1.POST("/login", u.AuthenticateUser)
	v1.PUT("/users/:id/password", u.UpdateUserPassword)
	v1.DELETE("/users/:id", u.DeleteUser)
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

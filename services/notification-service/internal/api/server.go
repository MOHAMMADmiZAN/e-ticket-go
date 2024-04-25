package api

import (
	"notification-service/docs"
	"notification-service/internal/api/handler"
	"notification-service/internal/api/middleware"
	"notification-service/internal/config"
	"notification-service/internal/repository"
	"notification-service/internal/services"

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
	swaggerDocConfig()
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
	v1 := s.Router.Group("/api/v1/notifications")

	// Health check route
	s.setupHealthCheckRoute()

	// Define handlers
	n := handler.NewNotificationHandler(services.NewNotificationService(repository.NewNotificationRepository(s.DB.Conn)))
	h := handler.NewUserNotificationHandler(services.NewUserNotificationService(repository.NewUserNotificationRepository(s.DB.Conn)))

	// Notification routes
	s.setupNotificationRoutes(v1, n)

	// User Notification routes
	s.setUpUserNotificationRoutes(v1, h)

	// Catch-all route for handling unmatched routes (404 Not Found)
	s.setupNoRouteHandler()
}

func (s *Server) setupHealthCheckRoute() {
	s.Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Notification services is running"})
	})
}

func (s *Server) setupNotificationRoutes(v1 *gin.RouterGroup, n *handler.NotificationHandler) {
	v1.POST("/", n.CreateNotification)
	v1.GET("/:id", n.GetNotification)
	v1.GET("/", n.ListNotifications)
	v1.DELETE("/:id", n.DeleteNotification)
}

func (s *Server) setUpUserNotificationRoutes(v1 *gin.RouterGroup, h *handler.UserNotificationHandler) {
	v1.GET("/users/:userID/preferences", h.GetUserPreferences)
	v1.POST("/users/:userID/preferences", h.CreateUserPreferences)
	v1.PUT("/users/:userID/preferences", h.UpdateUserPreferences)
	v1.DELETE("/users/:userID/preferences", h.DeleteUserPreferences)

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

func swaggerDocConfig() {
	host := os.Getenv("HOST")
	if host == "" {
		log.Fatal("HOST environment variable not set")
	}
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = "/api/v1/notifications"
	docs.SwaggerInfo.Title = "Notification Service API"
	docs.SwaggerInfo.Description = "Provides comprehensive endpoints for managing notifications, including create, update, and delete operations."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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

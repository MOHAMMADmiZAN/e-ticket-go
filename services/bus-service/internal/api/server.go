package api

import (
	_ "bus-service/docs" // Required for Swagger docs
	"bus-service/internal/api/handler"
	"bus-service/internal/api/middleware"
	"bus-service/internal/config"
	"bus-service/internal/repository"
	"bus-service/internal/services"
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

	b := handler.NewBusHandler(services.NewBusService(repository.NewBusRepository(s.DB.Conn)))

	// API Versioning
	v1 := s.Router.Group("/api/v1")

	// Health check route
	s.setupHealthCheckRoute()

	// Setup bus handlers
	s.setupBusRoutes(v1, b)

	// Catch-all route for handling unmatched routes (404 Not Found)
	s.setupNoRouteHandler()
}

func (s *Server) setupHealthCheckRoute() {
	s.Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Bus services is running"})
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

// setup Bus routes
func (s *Server) setupBusRoutes(v1 *gin.RouterGroup, b *handler.BusHandler) {
	//bus routes
	busGroup := v1.Group("/buses")
	{
		busGroup.GET("/", b.GetAllBuses)
		busGroup.POST("/", b.CreateBus)
		busGroup.GET("/:id", b.GetBusByID)
		busGroup.PUT("/:id", b.UpdateBus)
		busGroup.DELETE("/:id", b.DeleteBus)
		busGroup.GET("/status", b.GetBusesByStatus)
		busGroup.PUT("/:id/service-dates", b.UpdateBusServiceDates)
	}

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

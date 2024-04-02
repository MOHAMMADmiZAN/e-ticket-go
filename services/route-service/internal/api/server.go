package api

import (
	"context"
	"e-ticket/services/route-service/internal/api/handler"
	"e-ticket/services/route-service/internal/config"
	"e-ticket/services/route-service/internal/repository"
	"e-ticket/services/route-service/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
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
	r.Use(gin.Recovery(), gin.Logger()) // Add Logger middleware

	s := &Server{
		Router: r,
		DB:     databaseClient,
	}
	s.routes()
	return s
}

// routes registers all the routes to the router.
func (s *Server) routes() {
	// Create a new route handler and pass the route service.
	rh := handler.NewRouteHandler(service.NewRouteService(repository.NewRouteRepository(s.DB.Conn)))

	//Routes for the API
	v1 := s.Router.Group("/api/v1")
	{
		routes := v1.Group("/routes")
		{
			routes.POST("/", rh.CreateRoute)
			routes.GET("/", rh.GetAllRoutes)
			routes.GET("/:id", rh.GetRouteByID)
			routes.GET("health", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
			})
			//routes.PUT("/:id", rh.UpdateRoute)
			//routes.DELETE("/:id", rh.DeleteRoute)
		}
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
		log.Printf("Server started on %s\n", addr)
	}()

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

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
	// Define handlers
	rh := handler.NewRouteHandler(service.NewRouteService(repository.NewRouteRepository(s.DB.Conn)))
	sh := handler.NewStopHandler(service.NewStopService(repository.NewStopRepository(s.DB.Conn)))
	sch := handler.NewScheduleHandler(service.NewScheduleService(repository.NewScheduleRepository(s.DB.Conn), 5*time.Minute))

	// API Versioning
	v1 := s.Router.Group("/api/v1")

	// Health check route
	s.setupHealthCheckRoute()

	// Setup route groups
	s.setupRouteHandlers(v1, rh, sh, sch)

	// Catch-all route for handling unmatched routes (404 Not Found)
	s.setupNoRouteHandler()
}

func (s *Server) setupHealthCheckRoute() {
	s.Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
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

func (s *Server) setupRouteHandlers(v1 *gin.RouterGroup, rh *handler.RouteHandler, sh handler.StopHandlerInterface, sch *handler.ScheduleHandler) {
	routesGroup := v1.Group("/routes")
	// Route handlers
	{
		routesGroup.POST("/", rh.CreateRoute)
		routesGroup.GET("/", rh.GetAllRoutes)
		routesGroup.GET("/:id", rh.GetRouteByID)
		routesGroup.DELETE("/:id", rh.DeleteRoute)
	}

	// Stop handlers
	s.setupStopHandlers(routesGroup, sh)

	// Schedule handlers
	s.setupScheduleHandlers(routesGroup, sch)
}

func (s *Server) setupStopHandlers(routesGroup *gin.RouterGroup, sh handler.StopHandlerInterface) {
	stopsGroup := routesGroup.Group("/:id/stops")
	{
		stopsGroup.GET("/", sh.ListStopsForRoute)
		stopsGroup.POST("/", sh.AddStopToRoute)
		stopsGroup.PUT("/:stopId", sh.UpdateStop)
		stopsGroup.DELETE("/:stopId", sh.DeleteStop)
	}
}

func (s *Server) setupScheduleHandlers(routesGroup *gin.RouterGroup, sch *handler.ScheduleHandler) {
	schedulesGroup := routesGroup.Group("/:id/schedules")
	{
		schedulesGroup.POST("/", sch.CreateSchedule)
		schedulesGroup.GET("/", sch.GetSchedulesByRouteID)
		schedulesGroup.GET("/:scheduleID", sch.GetScheduleByID)
		schedulesGroup.PUT("/:scheduleID", sch.UpdateSchedule)
		schedulesGroup.DELETE("/:scheduleID", sch.DeleteSchedule)
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

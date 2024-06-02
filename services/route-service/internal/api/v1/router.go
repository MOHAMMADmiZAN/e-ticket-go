package v1

//	@title			My Route Service API
//	@version		1.0
//	@description	This API serves as an interface to interact with the My Route Service platform, providing endpoints for managing bus routes, Stops, and schedules.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Mohammad Mizan
//	@contact.url	http://swagger.io/support
//	@contact.email	takbir.jcd@gmail.com

//	@license.name	Apache License Version 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/api/v1

import (
	"net/http"
	"os"
	v1docs "route-service/internal/api/v1/docs"
	"route-service/internal/api/v1/handler"
	"route-service/internal/config"
	"route-service/internal/repository"
	"route-service/internal/services"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes registers all the routes to the router.
func SetupRoutes(router *gin.Engine, db *config.Database) {

	// Retrieve the Swagger host from environment variables
	swaggerHost := os.Getenv("SWAGGER_HOST")
	if swaggerHost == "" {
		swaggerHost = "localhost:8081" // default to localhost if not set
	}

	// Set the Swagger host
	v1docs.SwaggerInfo.Schemes = []string{"http", "https"}
	v1docs.SwaggerInfo.Host = swaggerHost
	swagUrl := ginSwagger.URL("/swagger/v1/doc.json")

	// v1 Swagger documentation
	router.GET("/swagger/v1/*any", func(c *gin.Context) {
		ginSwagger.WrapHandler(swaggerFiles.Handler, swagUrl)(c)
	})
	// Define handlers
	rh := v1.NewRouteHandler(services.NewRouteService(repository.NewRouteRepository(db.Conn)))
	sh := v1.NewStopHandler(services.NewStopService(repository.NewStopRepository(db.Conn)))
	sch := v1.NewScheduleHandler(services.NewScheduleService(repository.NewScheduleRepository(db.Conn), 5*time.Minute))

	// API Versioning
	v1Group := router.Group("/api/v1/routes")

	// Health check route
	setupHealthCheckRoute(router)

	// Setup route groups
	setupRouteHandlers(v1Group, rh)

	// Setup stop handlers
	setupStopHandlers(v1Group, sh)

	// Setup schedule handlers
	setupScheduleHandlers(v1Group, sch)

	// Catch-all route for handling unmatched routes (404 Not Found)
	setupNoRouteHandler(router)
}

func setupHealthCheckRoute(router *gin.Engine) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Route services is running"})
	})
}

func setupNoRouteHandler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
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

func setupRouteHandlers(rg *gin.RouterGroup, rh *v1.RouteHandler) {
	routesGroup := rg.Group("/")
	// Route handlers
	{
		routesGroup.POST("/", rh.CreateRoute)
		routesGroup.GET("/", rh.GetAllRoutes)
		routesGroup.GET("/:routeId", rh.GetRouteByID)
		routesGroup.PUT("/:routeId", rh.UpdateRoute)
		routesGroup.DELETE("/:routeId", rh.DeleteRoute)
	}
}

func setupStopHandlers(rg *gin.RouterGroup, sh v1.StopHandlerInterface) {
	stopsGroup := rg.Group("/:routeId/stops")
	// Stop handlers
	{
		stopsGroup.GET("/", sh.ListStopsForRoute)
		stopsGroup.POST("/", sh.AddStopToRoute)
		stopsGroup.GET("/:id", sh.GetStopByID)
		stopsGroup.PUT("/:id", sh.UpdateStop)
		stopsGroup.DELETE("/:id", sh.DeleteStop)
	}
}

func setupScheduleHandlers(rg *gin.RouterGroup, sch *v1.ScheduleHandler) {
	schedulesGroup := rg.Group("/stops/:stopId/schedules")
	{
		schedulesGroup.POST("/", sch.CreateSchedule)
		schedulesGroup.GET("/", sch.GetSchedules)
		schedulesGroup.GET("/:id", sch.GetScheduleByID)
		schedulesGroup.PUT("/:id", sch.UpdateSchedule)
		schedulesGroup.DELETE("/:id", sch.DeleteSchedule)
	}
}

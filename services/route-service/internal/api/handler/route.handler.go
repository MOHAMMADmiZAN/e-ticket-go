package handler

import (
	"e-ticket/services/route-service/internal/api/dto"
	"e-ticket/services/route-service/internal/service"
	"e-ticket/services/route-service/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RouteHandler holds the route service for handling requests.
type RouteHandler struct {
	routeService *service.RouteService
}

// NewRouteHandler creates a new handler for route endpoints.
func NewRouteHandler(routeService *service.RouteService) *RouteHandler {
	return &RouteHandler{
		routeService: routeService,
	}
}

// CreateRoute handles POST requests to create a new route.
func (h *RouteHandler) CreateRoute(c *gin.Context) {
	var jsonRequest dto.RouteCreateRequest
	if err := c.ShouldBindJSON(&jsonRequest); err != nil {
		validationErrors := util.FormatValidationError(err, jsonRequest)
		errResp := util.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Errors:  validationErrors,
		}
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	createdRoute, err := h.routeService.CreateRoute(c, jsonRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create route"})
		return
	}

	c.JSON(http.StatusCreated, createdRoute)
}

// GetAllRoutes handles GET requests to retrieve all routes.
func (h *RouteHandler) GetAllRoutes(c *gin.Context) {
	routes, err := h.routeService.GetRoutes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve routes"})
		return
	}

	c.JSON(http.StatusOK, routes)
}

// GetRouteByID handles GET requests to retrieve a route by ID.
func (h *RouteHandler) GetRouteByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		return
	}

	route, err := h.routeService.GetRouteByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve route"})
		return
	}

	c.JSON(http.StatusOK, route)
}

//// UpdateRoute handles PUT requests to update a route.
//func (h *RouteHandler) UpdateRoute(c *gin.Context) {
//	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
//		return
//	}
//
//	var route model.Route
//	if err := c.ShouldBindJSON(&route); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	updatedRoute, err := h.routeService.UpdateRoute(c, uint(id), route.Name, route.StartTime, route.Duration, route.StartLocation, route.EndLocation)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update route"})
//		return
//	}
//
//	c.JSON(http.StatusOK, updatedRoute)
//}

// DeleteRoute handles DELETE requests to delete a route.
func (h *RouteHandler) DeleteRoute(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID"})
		return
	}

	if err := h.routeService.DeleteRoute(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete route"})
		return
	}

	c.Status(http.StatusOK)
}

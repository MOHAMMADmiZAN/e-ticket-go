package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"route-service/internal/api/dto"
	"route-service/internal/services"
	"route-service/pkg"
	"strconv"
)

// RouteHandler holds the route services for handling requests.
type RouteHandler struct {
	routeService *services.RouteService
}

// NewRouteHandler creates a new handler for route endpoints.
func NewRouteHandler(routeService *services.RouteService) *RouteHandler {
	return &RouteHandler{
		routeService: routeService,
	}
}

// CreateRoute adds a new route to the system.
// @Summary Add a new route
// @Description Create a new route with the specified details.
// @Tags routes
// @Accept json
// @Produce json
// @Param route body dto.RouteCreateRequest true "Create Route Request"
// @Success 201 {object} dto.RouteInfo "Successfully created route information"
// @Failure 404 {string} string "Route not found"
// @Failure 500 {string} string "Unable to create route"
// @Router  / [post]
func (h *RouteHandler) CreateRoute(c *gin.Context) {
	var jsonRequest dto.RouteCreateRequest

	// Bind the JSON request to the dto.RouteCreateRequest struct
	if err := c.ShouldBindJSON(&jsonRequest); err != nil {
		// Log the error for internal tracking.
		log.Printf("Error binding JSON to RouteCreateRequest: %v", err)

		//Format the validation errors for the response.
		validationErrors := pkg.FormatValidationError(err, jsonRequest)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})
		return
	}

	// Convert the DTO to the Route models.
	route := jsonRequest.ToModel()

	// Attempt to create the route using the services layer.
	createdRoute, err := h.routeService.CreateRoute(c, route)
	if err != nil {
		// Log the error for internal tracking.
		log.Printf("Error creating route: %v", err)

		// Check if the error is a known type of error, or respond with a generic error message.
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.JSON(http.StatusNotFound, pkg.NewErrorResponse("Route not found"))
		case errors.Is(err, gorm.ErrInvalidData):
			c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid route data"))
		default:
			c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse("Unable to create route"))
		}
		return
	}

	// Convert the created route models back to a DTO and respond with it.
	routeInfo := dto.RouteModelToRouteInfo(createdRoute)
	c.JSON(http.StatusCreated, routeInfo)
}

// GetAllRoutes godoc
// @Summary Retrieve all routes
// @Description Get a list of all routes available in the system
// @Tags routes
// @Accept json
// @Produce json
// @Success 200 {array} dto.RouteInfo "List of all routes"
// @Failure 500 {object} pkg.ErrorMessage "Unable to retrieve routes"
// @Router  / [get]
// GetAllRoutes handles GET requests to retrieve all routes.
func (h *RouteHandler) GetAllRoutes(c *gin.Context) {
	routes, err := h.routeService.GetRoutes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse("Unable to retrieve routes"))
		return
	}

	c.JSON(http.StatusOK, routes)
}

// GetRouteByID godoc
// @Summary Get route by ID
// @Description Retrieve a route by its unique ID
// @Tags routes
// @Accept json
// @Produce json
//
//	@Param routeId path int true "Route ID"
//
// @Success 200 {object} dto.RouteResponse "Successfully retrieved route"
// @Failure 400 {object} pkg.ErrorMessage "Invalid route ID"
// @Failure 404 {object} pkg.ErrorMessage "Route not found"
// @Failure 500 {object} pkg.ErrorMessage "Unable to retrieve route"
// @Router  /{routeId} [get]
func (h *RouteHandler) GetRouteByID(c *gin.Context) {
	// Extract and parse route ID from the request parameters
	idParam := c.Param("routeId")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		// Log the error with the invalid ID for debugging
		log.Printf("Error parsing route ID '%s': %v", idParam, err)
		// Respond with a user-friendly error message
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid route ID"))
		return
	}

	// Retrieve the route by calling the services layer
	route, err := h.routeService.GetRouteByID(c, uint(id))
	if err != nil {
		// Determine if the error is a not-found error or a more serious error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Respond with not found if the route doesn't exist
			c.JSON(http.StatusNotFound, pkg.NewErrorResponse("Route not found"))
		} else {
			// Log the error for internal diagnostics
			log.Printf("Unable to retrieve route with ID %d: %v", id, err)
			// Respond with a generic error message
			c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse("Unable to retrieve route"))
		}
		return
	}

	// Successfully retrieved the route, respond with the route data
	c.JSON(http.StatusOK, route)
}

// UpdateRoute godoc
// @Summary Update a route
// @Description Update the details of an existing route by its ID.
// @Tags routes
// @Accept json
// @Produce json
//
//	@Param routeId path int true "Route ID" Format(uint)
//
// @Param route body dto.RouteUpdateRequest true "Route Update Information"
// @Success 200 {object} dto.RouteInfo "Successfully updated route information"
// @Failure 400 {object} pkg.ErrorMessage "Invalid route ID or request format"
// @Failure 500 {object} pkg.ErrorMessage "Internal Server Error - Unable to update route"
// @Router  /{routeId} [put]
func (h *RouteHandler) UpdateRoute(c *gin.Context) {

	// Parse the route ID from the request URI
	idParam := c.Param("routeId")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		// Log the error for debugging purposes
		log.Printf("Error parsing route ID '%s': %v", idParam, err)
		// Respond with a user-friendly error message
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid route ID"))
		return
	}
	var request dto.RouteUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {

		//Format the validation errors for the response.
		validationErrors := pkg.FormatValidationError(err, request)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})
		return
	}
	// CHECK IF THE ROUTE EXISTS
	existingRoute, err := h.routeService.GetRouteByID(c, uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, pkg.NewErrorResponse("Route not found"))
			return
		}
		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse("Unable to update route"))
		return
	}
	//if found marge with update request

	updatedRoute, err := h.routeService.UpdateRoute(c, request.ToModel(existingRoute))
	if err != nil {
		// Here you would handle different types of errors appropriately
		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse("Unable to update route"))
		return
	}

	c.JSON(http.StatusOK, updatedRoute)
}

// DeleteRoute godoc
// @Summary Delete a route
// @Description Delete a route by its unique ID
// @Tags routes
// @Accept json
// @Produce json
//
//	@Param routeId path int true "Route ID"
//
// @Success 204 "Route successfully deleted"
// @Failure 400 {object} pkg.ErrorMessage "Invalid route ID"
// @Failure 500 {object} pkg.ErrorMessage "Unable to delete route"
// @Router  /{routeId} [delete]
func (h *RouteHandler) DeleteRoute(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("routeId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid route ID"))
		return
	}

	if err := h.routeService.DeleteRoute(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse("Unable to delete route"))
		return
	}

	c.Status(http.StatusNoContent)
}

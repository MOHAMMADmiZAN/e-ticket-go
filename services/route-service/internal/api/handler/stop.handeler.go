package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"route-service/internal/api/dto"
	"route-service/internal/model"
	"route-service/internal/service"
	"strconv"
)

type StopHandlerInterface interface {
	AddStopToRoute(c *gin.Context)
	ListStopsForRoute(c *gin.Context)
	UpdateStop(c *gin.Context)
	DeleteStop(c *gin.Context)
}

type StopHandler struct {
	stopService service.StopServiceInterface
}

// Ensure StopHandler implements the StopHandlerInterface.
var _ StopHandlerInterface = &StopHandler{}

func NewStopHandler(stopService service.StopServiceInterface) StopHandlerInterface {
	return &StopHandler{
		stopService: stopService,
	}
}

func (h *StopHandler) AddStopToRoute(c *gin.Context) {
	var req dto.AddStopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	routeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid route ID", "details": err.Error()})
		return
	}

	stop := model.Stop{
		RouteID:  uint(routeID),
		Name:     req.Name,
		Sequence: req.Sequence,
	}

	createdStop, err := h.stopService.AddStopToRoute(c.Request.Context(), stop.RouteID, stop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not add stop", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdStop)
}

// ListStopsForRoute handles HTTP GET requests to list all stops for a given route.
func (h *StopHandler) ListStopsForRoute(c *gin.Context) {
	// Extract the routeID from the URL parameter and validate it.
	routeIDParam := c.Param("id")
	routeID, err := strconv.ParseUint(routeIDParam, 10, 32)
	if err != nil {
		// Respond with an error if the route ID is not a valid unsigned integer.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid route ID format", "details": err.Error()})
		return
	}

	// Call the service layer to get the list of stops for the given route.
	stops, err := h.stopService.GetStopsByRouteID(c.Request.Context(), uint(routeID))
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve stops", "details": err.Error()})
		return
	}

	// If the stops are retrieved successfully, respond with the list of stops.
	c.JSON(http.StatusOK, stops)
}

func (h *StopHandler) UpdateStop(c *gin.Context) {
	routeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	stopID, err := strconv.ParseUint(c.Param("stopId"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID", "details": err.Error()})
		return
	}

	var req dto.UpdateStopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	stop := model.Stop{
		ID:       uint(stopID),
		RouteID:  uint(routeID),
		Name:     req.Name,
		Sequence: req.Sequence,
	}

	updatedStop, err := h.stopService.UpdateStop(c.Request.Context(), stop.RouteID, stop.ID, stop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update stop", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedStop)
}

func (h *StopHandler) DeleteStop(c *gin.Context) {
	routeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	stopID, err := strconv.ParseUint(c.Param("stopId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID", "details": err.Error()})
		return
	}

	err = h.stopService.DeleteStop(c.Request.Context(), uint(routeID), uint(stopID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete stop", "details": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

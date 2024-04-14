package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"route-service/internal/api/dto"
	"route-service/internal/service"
	"strconv"
)

// ScheduleHandler holds the necessary components for the related handler functions.
type ScheduleHandler struct {
	scheduleService service.ScheduleService
}

// NewScheduleHandler creates a new instance of ScheduleHandler.
func NewScheduleHandler(scheduleService service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleService: scheduleService,
	}
}

// CreateSchedule handles POST requests to create a new schedule.
func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	// Parse the route ID from the URL parameter
	routeIDParam := c.Param("id")
	routeID, err := strconv.ParseUint(routeIDParam, 10, 32) // Convert string to uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Route ID in URL"})
		return
	}
	var req dto.AddScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if uint(routeID) != req.RouteID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Route ID in URL does not match Route ID in request body"})
		return
	}

	// Convert DTO to Model and validate the request
	newSchedule := req.ToModel()

	resp, err := h.scheduleService.CreateSchedule(c.Request.Context(), *newSchedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetScheduleByID handles GET requests to retrieve a schedule by its ID.
func (h *ScheduleHandler) GetScheduleByID(c *gin.Context) {
	scheduleID, err := strconv.ParseUint(c.Param("scheduleID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid schedule ID"})
		return
	}

	resp, err := h.scheduleService.GetScheduleByID(c.Request.Context(), uint(scheduleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetSchedulesByRouteID handles GET requests to retrieve all schedules for a route.
func (h *ScheduleHandler) GetSchedulesByRouteID(c *gin.Context) {
	routeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid route ID"})
		return
	}

	resp, err := h.scheduleService.GetSchedulesByRouteID(c.Request.Context(), uint(routeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateSchedule handles PUT requests to update a schedule.
func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	scheduleID, err := strconv.ParseUint(c.Param("scheduleID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid schedule ID"})
		return
	}

	var req dto.UpdateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	// Convert DTO to Model and validate the request
	updateSchedule := req.ToModel()

	resp, err := h.scheduleService.UpdateSchedule(c.Request.Context(), uint(scheduleID), *updateSchedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteSchedule handles DELETE requests to remove a schedule.
func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	scheduleID, err := strconv.ParseUint(c.Param("scheduleID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid schedule ID"})
		return
	}

	if err := h.scheduleService.DeleteSchedule(c.Request.Context(), uint(scheduleID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

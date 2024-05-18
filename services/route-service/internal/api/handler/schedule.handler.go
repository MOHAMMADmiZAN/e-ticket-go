package handler

import (
	"net/http"
	"route-service/internal/api/dto"
	"route-service/internal/services"
	"route-service/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ScheduleHandler holds the necessary components for the related handler functions.
type ScheduleHandler struct {
	scheduleService services.ScheduleService
}

// NewScheduleHandler creates a new instance of ScheduleHandler.
func NewScheduleHandler(scheduleService services.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		scheduleService: scheduleService,
	}
}

// CreateSchedule @Summary Create Schedule
// @Description Create a new schedule for a specific route.
// @Tags Schedules
// @Accept json
// @Produce json
// @Param stopId path int true "Stop ID" minimum(1)
// @Param addScheduleRequest body dto.AddScheduleRequest true "Schedule information"
// @Success 201 {object} dto.ScheduleResponse "Created schedule details"
// @Failure 400 {object} pkg.ErrorMessage "Invalid route ID, request format, or validation error"
// @Failure 404 {object} pkg.ErrorMessage "Route not found"
// @Failure 500 {object} pkg.ErrorMessage "Internal server error"
// @Router /stops/{stopId}/schedules [post]
func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	// Parse the route ID from the URL parameter
	stopIdPrams := c.Param("stopId")
	stopId, err := strconv.ParseUint(stopIdPrams, 10, 32) // Convert string to uint
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid Stop ID"))
		return
	}
	var req dto.AddScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := pkg.FormatValidationError(err, req)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse(err.Error()))
		return
	}

	if uint(stopId) != req.StopID {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid stop ID"))
		return
	}

	// Convert DTO to Model and validate the request
	newSchedule := req.ToModel()

	resp, err := h.scheduleService.CreateSchedule(c.Request.Context(), *newSchedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetScheduleByID @Summary Get Schedule by ID
// @Description Retrieve a schedule by its ID.
// @Tags Schedules
// @Accept json
// @Produce json
// @Param stopId path int true "Stop ID" minimum(1)
// @Param id path int true "Schedule ID" minimum(1)
// @Success 200 {object} dto.ScheduleResponse "Schedule details"
// @Failure 400 {object} pkg.ErrorMessage "Invalid schedule ID"
// @Failure 404 {object} pkg.ErrorMessage "Schedule not found"
// @Failure 500 {object} pkg.ErrorMessage "Internal server error"
// @Router /stops/{stopId}/schedules/{id} [get]
// GetScheduleByID handles GET requests to retrieve a schedule by its ID.
func (h *ScheduleHandler) GetScheduleByID(c *gin.Context) {
	scheduleID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	stopId, err1 := strconv.ParseUint(c.Param("stopId"), 10, 32)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid Stop ID"))
		return

	}
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid schedule ID"))
		return
	}

	resp, err := h.scheduleService.GetScheduleByID(c.Request.Context(), uint(scheduleID), uint(stopId))
	if err != nil {
		c.JSON(http.StatusNotFound, pkg.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetSchedulesByRouteID @Summary Get Schedules by Route ID
// @Description Retrieve all schedules for a route by its ID.
// @Tags Schedules
// @Accept json
// @Produce json
// @Param routeId path int true "Route ID" minimum(1)
// @Success 200 {array} dto.ScheduleResponse "List of schedules"
// @Failure 400 {object} pkg.ErrorMessage "Invalid route ID"
// @Failure 404 {object} pkg.ErrorMessage "Route not found"
// @Failure 500 {object} pkg.ErrorMessage "Internal server error"
// @Router /{routeId}/schedules [get]
// GetSchedulesByRouteID handles GET requests to retrieve all schedules for a route.
func (h *ScheduleHandler) GetSchedulesByRouteID(c *gin.Context) {
	routeID, err := strconv.ParseUint(c.Param("routeId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid route ID"))
		return
	}

	resp, err := h.scheduleService.GetSchedulesByRouteID(c.Request.Context(), uint(routeID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateSchedule @Summary Update Schedule
// @Description Update details of a schedule by its ID.
// @Tags Schedules
// @Accept json
// @Produce json
// @Param routeId path int true "Route ID" minimum(1)
// @Param id path int true "Schedule ID" minimum(1)
// @Param updateScheduleRequest body dto.UpdateScheduleRequest true "Updated schedule information"
// @Success 200 {object} dto.ScheduleResponse "Updated schedule details"
// @Failure 400 {object} pkg.ErrorMessage "Invalid schedule ID or request format"
// @Failure 404 {object} pkg.ErrorMessage "Schedule not found"
// @Failure 500 {object} pkg.ErrorMessage "Internal server error"
// @Router /{routeId}/schedules/{id} [put]
// UpdateSchedule handles PUT requests to update a schedule.
func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	scheduleID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid schedule ID"))
		return
	}

	var req dto.UpdateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := pkg.FormatValidationError(err, req)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse(err.Error()))
		return

	}

	// Convert DTO to Model and validate the request
	updateSchedule := req.ToModel()

	resp, err := h.scheduleService.UpdateSchedule(c.Request.Context(), uint(scheduleID), *updateSchedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteSchedule @Summary Delete Schedule
// @Description Remove a schedule by its ID.
// @Tags Schedules
// @Accept json
// @Produce json
// @Param routeId path int true "Route ID" minimum(1)
// @Param id path int true "Schedule ID" minimum(1)
// @Success 204 "No Content"
// @Failure 400 {object} pkg.ErrorMessage "Invalid schedule ID"
// @Failure 404 {object} pkg.ErrorMessage "Schedule not found"
// @Failure 500 {object} pkg.ErrorMessage "Internal server error"
// @Router /{routeId}/schedules/{id} [delete]
// DeleteSchedule handles DELETE requests to remove a schedule.
func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	scheduleID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	routeId, err1 := strconv.ParseUint(c.Param("routeId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid schedule ID"))
		return
	}
	if err1 != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid route ID"))
		return

	}

	if err := h.scheduleService.DeleteSchedule(c.Request.Context(), uint(scheduleID), uint(routeId)); err != nil {
		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}

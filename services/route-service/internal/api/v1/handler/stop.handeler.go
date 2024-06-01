package v1

import (
	"net/http"
	"route-service/internal/api/dto"
	"route-service/internal/models"
	"route-service/internal/services"
	"route-service/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StopHandlerInterface interface {
	AddStopToRoute(c *gin.Context)
	ListStopsForRoute(c *gin.Context)
	UpdateStop(c *gin.Context)
	DeleteStop(c *gin.Context)
	GetStopByID(c *gin.Context)
}

type StopHandler struct {
	stopService services.StopServiceInterface
}

// Ensure StopHandler implements the StopHandlerInterface.
var _ StopHandlerInterface = &StopHandler{}

func NewStopHandler(stopService services.StopServiceInterface) StopHandlerInterface {
	return &StopHandler{
		stopService: stopService,
	}
}

// AddStopToRoute @Summary Add Stop to Route
//
//	@Description	Add a new stop to the specified route.
//	@Tags			Stops
//	@Accept			json
//	@Produce		json
//	@Param			routeId	path		int					true	"Route ID"	minimum(1)
//	@Param			stop	body		dto.AddStopRequest	true	"Stop information"
//	@Success		201		{object}	dto.StopResponse	"createdStop"
//	@Failure		400		{object}	pkg.ErrorMessage	"Invalid request or incorrect data"
//	@Failure		404		{object}	pkg.ErrorMessage	"Route not found"
//	@Failure		500		{object}	pkg.ErrorMessage	"Internal server error"
//	@Router			/routes/{routeId}/stops [post]
func (h *StopHandler) AddStopToRoute(c *gin.Context) {
	var req dto.AddStopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//Format the validation errors for the response.
		validationErrors := pkg.FormatValidationError(err, req)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})
		return
	}

	routeID, err := strconv.ParseUint(c.Param("routeId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid route ID"))
		return
	}

	stop := models.Stop{
		RouteID:  uint(routeID),
		Name:     req.Name,
		Sequence: req.Sequence,
	}

	createdStop, err := h.stopService.AddStopToRoute(c.Request.Context(), stop.RouteID, stop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, createdStop)
}

// ListStopsForRoute @Summary List Stops for Route
//
//	@Description	Retrieve a list of stops for a given route.
//	@Tags			Stops
//	@Accept			json
//	@Produce		json
//	@Param			routeId	path		int					true	"Route ID"	minimum(1)
//	@Success		200		{array}		dto.StopResponse	"stops"
//	@Failure		400		{object}	pkg.ErrorMessage	"Invalid route ID"
//	@Failure		404		{object}	pkg.ErrorMessage	"Route not found"
//	@Failure		500		{object}	pkg.ErrorMessage	"Internal server error"
//	@Router			/routes/{routeId}/stops [get]
//
// ListStopsForRoute handles HTTP GET requests to list all stops for a given route.
func (h *StopHandler) ListStopsForRoute(c *gin.Context) {
	// Extract the routeID from the URL parameter and validate it.
	routeIDParam := c.Param("routeId")
	routeID, err := strconv.ParseUint(routeIDParam, 10, 32)
	if err != nil {
		// Respond with an error if the route ID is not a valid unsigned integer.
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid route ID"))
		return
	}

	// Call the services layer to get the list of stops for the given route.
	stops, err := h.stopService.GetStopsByRouteID(c.Request.Context(), uint(routeID))
	if err != nil {

		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse(err.Error()))
		return
	}

	// If the stops are retrieved successfully, respond with the list of stops.
	c.JSON(http.StatusOK, stops)
}

// UpdateStop @Summary Update Stop
//
//	@Description	Update details of a stop for a given route.
//	@Tags			Stops
//	@Accept			json
//	@Produce		json
//	@Param			routeId				path		int						true	"Route ID"	minimum(1)
//	@Param			id					path		int						true	"Stop ID"	minimum(1)
//	@Param			updateStopRequest	body		dto.UpdateStopRequest	true	"Updated stop information"
//	@Success		200					{object}	dto.StopResponse		"updatedStop"
//	@Failure		400					{object}	pkg.ErrorMessage		"Invalid route ID or stop ID"
//	@Failure		404					{object}	pkg.ErrorMessage		"Route or stop not found"
//	@Failure		500					{object}	pkg.ErrorMessage		"Internal server error"
//	@Router			/routes/{routeId}/stops/{id} [put]
func (h *StopHandler) UpdateStop(c *gin.Context) {
	routeID, err := strconv.ParseUint(c.Param("routeId"), 10, 32)
	stopID, err1 := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid route ID"))
		return
	}
	if err1 != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid stop ID"))
		return

	}

	var req dto.UpdateStopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//Format the validation errors for the response.
		validationErrors := pkg.FormatValidationError(err, req)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})
		return
	}

	stop := models.Stop{
		ID:       uint(stopID),
		RouteID:  uint(routeID),
		Name:     req.Name,
		Sequence: req.Sequence,
	}

	updatedStop, err := h.stopService.UpdateStop(c.Request.Context(), stop.RouteID, stop.ID, stop)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, updatedStop)
}

// DeleteStop @Summary Delete Stop
//
//	@Description	Delete a stop from a given route.
//	@Tags			Stops
//	@Accept			json
//	@Produce		json
//	@Param			routeId	path	int	true	"Route ID"	minimum(1)
//	@Param			id		path	int	true	"Stop ID"	minimum(1)
//	@Success		200		"OK"
//	@Failure		400		{object}	pkg.ErrorMessage	"Invalid route ID or stop ID"
//	@Failure		404		{object}	pkg.ErrorMessage	"Route or stop not found"
//	@Failure		500		{object}	pkg.ErrorMessage	"Internal server error"
//	@Router			/routes/{routeId}/stops/{id} [delete]
func (h *StopHandler) DeleteStop(c *gin.Context) {
	routeID, err := strconv.ParseUint(c.Param("routeId"), 10, 32)
	stopID, err1 := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid route ID"))
		return
	}
	if err1 != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid stop ID"))
		return

	}

	err = h.stopService.DeleteStop(c.Request.Context(), uint(routeID), uint(stopID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

// GetStopByID @Summary Get Stop by ID
//
//	@Description	Retrieve details of a stop by its unique ID.
//	@Tags			Stops
//	@Accept			json
//	@Produce		json
//	@Param			routeId	path		int					true	"Route ID"	minimum(1)
//	@Param			id		path		int					true	"Stop ID"	minimum(1)
//	@Success		200		{object}	dto.StopResponse	"stop"
//	@Failure		400		{object}	pkg.ErrorMessage	"Invalid route ID or stop ID"
//	@Failure		404		{object}	pkg.ErrorMessage	"Route or stop not found"
//	@Failure		500		{object}	pkg.ErrorMessage	"Internal server error"
//	@Router			/routes/{routeId}/stops/{id} [get]
func (h *StopHandler) GetStopByID(c *gin.Context) {
	routeID, err := strconv.ParseUint(c.Param("routeId"), 10, 32)
	stopID, err1 := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid route ID"))
		return
	}
	if err1 != nil {
		c.JSON(http.StatusBadRequest, pkg.NewErrorResponse("Invalid stop ID"))
		return

	}

	stop, err := h.stopService.GetStopByID(c.Request.Context(), uint(routeID), uint(stopID))
	if err != nil {
		c.JSON(http.StatusNotFound, pkg.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, stop)
}

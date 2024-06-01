package handler

import (
	"bus-service/internal/api/dto"
	"bus-service/internal/models"
	"bus-service/internal/services"
	"bus-service/pkg"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BusHandler struct {
	busService *services.BusService
}

func NewBusHandler(busService *services.BusService) *BusHandler {
	return &BusHandler{busService: busService}
}

// GetAllBuses godoc
// @Summary Get all buses
// @Description Retrieve a list of all buses from the bus service.
// @Tags buses
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.BusResponse "List of all buses"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error - Could not retrieve buses"
//
//	@Router / [get]
func (h *BusHandler) GetAllBuses(c *gin.Context) {
	buses, err := h.busService.GetAllBuses()
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}
	busesResponse := make([]dto.BusResponse, 0, len(buses))
	for _, bus := range buses {
		busesResponse = append(busesResponse, dto.FromModel(bus))

	}
	pkg.RespondWithSuccess(c, http.StatusOK, busesResponse, "")
}

// GetBusByID godoc
// @Summary Get bus by ID
// @Description Retrieve details of a specific bus by its ID.
// @Tags buses
// @Accept  json
// @Produce  json
//
//	@Param busId path int true "Bus ID"
//
// @Success 200 {object} dto.BusResponse "Successfully retrieved bus"
// @Failure 400 {object} pkg.APIResponse "Bad Request - Invalid ID format or Bus not found"
//
//	@Router /{busId} [get]
func (h *BusHandler) GetBusByID(c *gin.Context) {
	bus, err := h.getBusFromIDParam(c)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, dto.FromModel(*bus), "")
}

// CreateBus
// @Summary Create a new bus
// @Description Add a new bus to the system with the provided details.
// @Tags buses
// @Accept json
// @Produce json
// @Param bus body dto.CreateBusDTO true "Bus Creation Data"
// @Success 201 {object} dto.BusResponse "Successfully created bus"
// @Failure 400 {object} pkg.APIResponse  "Bad Request - Invalid input provided"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error - Unable to create bus"
//
//	@Router / [post]
func (h *BusHandler) CreateBus(c *gin.Context) {
	var createBusDTO dto.CreateBusDTO
	if err := c.ShouldBindJSON(&createBusDTO); err != nil {
		//Format the validation errors for the response.
		validationErrors := pkg.FormatValidationError(err, createBusDTO)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})

		return
	}

	// Additional custom validation if necessary
	if err := dto.ValidateCreateBusDTO(createBusDTO); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, err)
		log.Print(`error:2 `, err)
		return
	}
	bus := createBusDTO.ToModel()
	newBus, err := h.busService.CreateBus(bus)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}
	pkg.RespondWithSuccess(c, http.StatusCreated, newBus, "Bus created successfully")
}

// UpdateBus godoc
// @Summary Update bus details
// @Description Update the details of a specific bus by its ID.
// @Tags buses
// @Accept  json
// @Produce  json
//
//	@Param busId path int true "Bus ID"
//
// @Param bus body dto.UpdateBusDTO true "Bus Update Data"
// @Success 200 {object} dto.BusResponse "Bus updated successfully"
// @Failure 400 {object} pkg.APIResponse "Bad Request - Invalid ID or update data format"
// @Failure 404 {object} pkg.APIResponse "Not Found - Bus not found"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error - Could not update bus"
//
//	@Router /{busId} [put]
func (h *BusHandler) UpdateBus(c *gin.Context) {
	exitingBus, err := h.getBusFromIDParam(c)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	var req dto.UpdateBusDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		//Format the validation errors for the response.
		validationErrors := pkg.FormatValidationError(err, req)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})
		return
	}
	bus := req.ToModel(exitingBus)
	updatedBus, err := h.busService.UpdateBus(*bus)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, dto.FromModel(*updatedBus), "Bus updated successfully")
}

// DeleteBus godoc
// @Summary Delete a bus
// @Description Delete the bus with the specified ID.
// @Tags buses
// @Accept json
// @Produce json
//
//	@Param busId path int true "Bus ID"
//
// @Success 200 {object} pkg.APIResponse "Bus deleted successfully"
// @Failure 400 {object} pkg.APIResponse "Bad Request - Invalid bus ID"
// @Failure 404 {object} pkg.APIResponse "Not Found - Bus not found"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error - Unable to delete bus"
//
//	@Router /{busId} [delete]
func (h *BusHandler) DeleteBus(c *gin.Context) {
	bus, err := h.getBusFromIDParam(c)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	err = h.busService.DeleteBus(bus.ID)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}
	pkg.RespondWithSuccess(c, http.StatusOK, nil, "Bus deleted successfully")
}

// getBusFromIDParam is a helper method to retrieve the bus by ID from the URL parameter.
func (h *BusHandler) getBusFromIDParam(c *gin.Context) (*models.Bus, error) {
	id, err := h.parseUintParam(c, "busId")
	if err != nil {
		return nil, err
	}

	bus, err := h.busService.GetBusByID(id)
	if err != nil {
		return nil, err
	}

	return bus, nil
}

// parseUintParam is a helper method to parse unsigned integer parameters from the URL.
func (h *BusHandler) parseUintParam(c *gin.Context, paramName string) (uint, error) {
	paramValue := c.Param(paramName)
	id, err := strconv.ParseUint(paramValue, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid %s: %s", paramName, paramValue)
	}
	return uint(id), nil
}

// GetBusesByStatus retrieves all buses by their operational status using a query parameter.
// @Summary Get all buses by status
// @Description Retrieve a list of all buses with the specified status.
// @Tags buses
// @Accept  json
// @Produce  json
// @Param status query string true "status" Enums(active, maintenance, decommissioned) "Bus operational status"
// @Success 200 {array} dto.BusResponse "List of all buses"
// @Failure 400 {object} pkg.APIResponse "Bad Request - Invalid status"
//
//	@Router /status [get]
func (h *BusHandler) GetBusesByStatus(c *gin.Context) {
	// Retrieve the status from the query string
	status := c.Query("status")
	if status == "" {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("missing status query parameter"))
		return
	}

	buses, err := h.busService.GetBusesByStatus(status)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	busesResponse := make([]dto.BusResponse, 0, len(buses))
	for _, bus := range buses {
		busesResponse = append(busesResponse, dto.FromModel(bus))
	}

	pkg.RespondWithSuccess(c, http.StatusOK, busesResponse, "")
}

// UpdateBusServiceDates godoc
// @Summary Update bus service dates
// @Description Update the last and next service dates for a specific bus.
// @Tags buses
// @Accept  json
// @Produce  json
//
//	@Param busId path int true "Bus ID"
//
// @Param serviceDates body dto.UpdateBusServiceDatesDTO true "Service Dates"
// @Success 200 {object} pkg.APIResponse "Bus service dates updated successfully"
// @Failure 400 {object} pkg.APIResponse "Bad Request - Invalid bus ID or service dates"
//
//	@Router /{busId}/service-dates [put]
func (h *BusHandler) UpdateBusServiceDates(c *gin.Context) {
	var req dto.UpdateBusServiceDatesDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		//Format the validation errors for the response.
		validationErrors := pkg.FormatValidationError(err, req)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})
		return
	}

	// Extracting bus ID from the path parameter
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	if err := h.busService.UpdateBusServiceDates(uint(id), req.LastServiceDate, req.NextServiceDate); err != nil {

		pkg.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bus service dates updated successfully"})
}

// GetBusesByRouteID godoc
// @Summary Get all buses by route ID
// @Description Retrieve a list of all buses that operate on a specific route.
// @Tags buses
// @Accept  json
// @Produce  json
//
//	@Param routeId path int true "Route ID"
//
// @Success 200 {array} dto.BusResponse "List of all buses"
// @Failure 400 {object} pkg.APIResponse "Bad Request - Invalid route ID"
//
//	@Router /routes/{routeId} [get]
func (h *BusHandler) GetBusesByRouteID(c *gin.Context) {
	routeID, err := h.parseUintParam(c, "routeId")
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, err)
		return
	}

	buses, err := h.busService.GetBusesByRoute(routeID)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	busesResponse := make([]dto.BusResponse, 0, len(buses))
	for _, bus := range buses {
		busesResponse = append(busesResponse, dto.FromModel(bus))
	}

	pkg.RespondWithSuccess(c, http.StatusOK, busesResponse, "")
}

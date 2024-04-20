package handler

import (
	"bus-service/internal/api/dto"
	"bus-service/internal/models"
	"bus-service/internal/services"
	"fmt"
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

// APIResponse represents a standard API response structure.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// respondWithError handles error responses with a standard format.
func respondWithError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, APIResponse{Success: false, Error: err.Error()})
}

// respondWithSuccess handles success responses with a standard format.
func respondWithSuccess(c *gin.Context, statusCode int, data interface{}, message string) {
	response := APIResponse{Success: true, Data: data}
	if message != "" {
		response.Message = message
	}
	c.JSON(statusCode, response)
}

// GetAllBuses godoc
// @Summary Get all buses
// @Description Retrieve a list of all buses from the bus service.
// @Tags buses
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.BusResponse "List of all buses"
// @Failure 500 {object} APIResponse "Internal Server Error - Could not retrieve buses"
// @Router /buses [get]
func (h *BusHandler) GetAllBuses(c *gin.Context) {
	buses, err := h.busService.GetAllBuses()
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	busesResponse := make([]dto.BusResponse, 0, len(buses))
	for _, bus := range buses {
		busesResponse = append(busesResponse, dto.FromModel(bus))

	}
	respondWithSuccess(c, http.StatusOK, busesResponse, "")
}

// GetBusByID godoc
// @Summary Get bus by ID
// @Description Retrieve details of a specific bus by its ID.
// @Tags buses
// @Accept  json
// @Produce  json
// @Param id path int true "Bus ID"
// @Success 200 {object} dto.BusResponse "Successfully retrieved bus"
// @Failure 400 {object} APIResponse "Bad Request - Invalid ID format or Bus not found"
// @Router /buses/{id} [get]
func (h *BusHandler) GetBusByID(c *gin.Context) {
	bus, err := h.getBusFromIDParam(c)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}

	respondWithSuccess(c, http.StatusOK, dto.FromModel(*bus), "")
}

// CreateBus
// @Summary Create a new bus
// @Description Add a new bus to the system with the provided details.
// @Tags buses
// @Accept json
// @Produce json
// @Param bus body dto.CreateBusDTO true "Bus Creation Data"
// @Success 201 {object} dto.BusResponse "Successfully created bus"
// @Failure 400 {object} APIResponse  "Bad Request - Invalid input provided"
// @Failure 500 {object} APIResponse "Internal Server Error - Unable to create bus"
// @Router /buses [post]
func (h *BusHandler) CreateBus(c *gin.Context) {
	var createBusDTO dto.CreateBusDTO
	if err := c.ShouldBindJSON(&createBusDTO); err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}

	// Additional custom validation if necessary
	if err := dto.ValidateCreateBusDTO(createBusDTO); err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}
	bus := createBusDTO.ToModel()
	newBus, err := h.busService.CreateBus(bus)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithSuccess(c, http.StatusCreated, newBus, "Bus created successfully")
}

// UpdateBus godoc
// @Summary Update bus details
// @Description Update the details of a specific bus by its ID.
// @Tags buses
// @Accept  json
// @Produce  json
// @Param id path int true "Bus ID"
// @Param bus body dto.UpdateBusDTO true "Bus Update Data"
// @Success 200 {object} dto.BusResponse "Bus updated successfully"
// @Failure 400 {object} APIResponse "Bad Request - Invalid ID or update data format"
// @Failure 404 {object} APIResponse "Not Found - Bus not found"
// @Failure 500 {object} APIResponse "Internal Server Error - Could not update bus"
// @Router /buses/{id} [put]
func (h *BusHandler) UpdateBus(c *gin.Context) {
	bus, err := h.getBusFromIDParam(c)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}

	if err := c.ShouldBindJSON(&bus); err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}

	updatedBus, err := h.busService.UpdateBus(*bus)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithSuccess(c, http.StatusOK, updatedBus, "Bus updated successfully")
}

// DeleteBus godoc
// @Summary Delete a bus
// @Description Delete the bus with the specified ID.
// @Tags buses
// @Accept json
// @Produce json
// @Param id path int true "Bus ID"
// @Success 200 {object} APIResponse "Bus deleted successfully"
// @Failure 400 {object} APIResponse "Bad Request - Invalid bus ID"
// @Failure 404 {object} APIResponse "Not Found - Bus not found"
// @Failure 500 {object} APIResponse "Internal Server Error - Unable to delete bus"
// @Router /buses/{id} [delete]
func (h *BusHandler) DeleteBus(c *gin.Context) {
	bus, err := h.getBusFromIDParam(c)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, err)
		return
	}

	err = h.busService.DeleteBus(bus.ID)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, err)
		return
	}
	respondWithSuccess(c, http.StatusOK, nil, "Bus deleted successfully")
}

// getBusFromIDParam is a helper method to retrieve the bus by ID from the URL parameter.
func (h *BusHandler) getBusFromIDParam(c *gin.Context) (*models.Bus, error) {
	id, err := h.parseUintParam(c, "id")
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

package handler

import (
	"bus-service/internal/api/dto"
	"bus-service/internal/models"
	"bus-service/internal/services"
	"bus-service/pkg"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SeatHandler struct {
	seatService services.ISeatService
}

func NewSeatHandler(seatService services.ISeatService) *SeatHandler {
	return &SeatHandler{
		seatService: seatService,
	}
}

// GetSeatsByBus @Summary Retrieve seat inventory for a bus
// @Description Retrieves the seat inventory for a specific bus
// @Tags seats
// @Accept  json
// @Produce  json
// @Param busID path int true "Bus ID"
// @Success 200 {array} dto.SeatResponse "Array of Seat objects"
// @Failure 400 {object} pkg.APIResponse "Bad Request"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error"
// @Router /{busID}/seats/ [get]
func (h *SeatHandler) GetSeatsByBus(c *gin.Context) {
	busID, err := strconv.ParseUint(c.Param("busID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid bus ID format"))
		return
	}

	seats, err := h.seatService.GetSeatsByBus(uint(busID))
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, seats, "")
}

// CreateSeat @Summary Add a new seat to a bus
// @Description Adds a new seat to a specific bus
// @Tags seats
// @Accept  json
// @Produce  json
// @Param busID path int true "Bus ID"
// @Param seat body dto.CreateSeatRequest true "Seat Data"
// @Success 201 {object} dto.SeatResponse "Successfully created seat"
// @Failure 400 {object} pkg.APIResponse "Bad Request"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error"
// @Router /{busID}/seats/ [post]
func (h *SeatHandler) CreateSeat(c *gin.Context) {
	var seatDTO dto.CreateSeatRequest
	if err := c.ShouldBindJSON(&seatDTO); err != nil {
		//Format the validation errors for the response.
		validationErrors := pkg.FormatValidationError(err, seatDTO)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})
		return
	}

	if err := h.seatService.CreateSeat(seatDTO.BusID, seatDTO.ToModel()); err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusCreated, nil, "Seat created successfully")
}

// GetSeatByID @Summary Get details of a seat
// @Description Retrieves details of a specific seat
// @Tags seats
// @Accept  json
// @Produce  json
// @Param id path int true "Seat ID"
// @Success 200 {object} dto.SeatResponse "Successfully retrieved seat"
// @Failure 400 {object} pkg.APIResponse "Bad Request"
// @Failure 404 {object} pkg.APIResponse "Not Found"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error"
// @Router /{busID}/seats/{id} [get]
func (h *SeatHandler) GetSeatByID(c *gin.Context) {
	seatID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid seat ID format"))
		return
	}

	seat, err := h.seatService.GetSeat(uint(seatID))
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, seat, "")
}

// UpdateSeat @Summary Update details or status of a seat
// @Description Updates the details or status of a specific seat
// @Tags seats
// @Accept  json
// @Produce  json
// @Param id path int true "Seat ID"
// @Param seat body dto.UpdateSeatRequest true "Seat update data"
// @Success 200 {object} dto.SeatResponse "Successfully updated seat"
// @Failure 400 {object} pkg.APIResponse "Bad Request"
// @Failure 404 {object} pkg.APIResponse "Not Found"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error"
// @Router /{busID}/seats/{id} [put]
func (h *SeatHandler) UpdateSeat(c *gin.Context) {
	seatID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid seat ID format"))
		return
	}

	var updateSeatDTO dto.UpdateSeatRequest
	if err := c.ShouldBindJSON(&updateSeatDTO); err != nil {
		//Format the validation errors for the response.
		validationErrors := pkg.FormatValidationError(err, updateSeatDTO)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})
		return
	}
	//FETCH EXITING SEAT
	exitingSeats, err := h.seatService.GetSeat(uint(seatID))

	if err := h.seatService.UpdateSeat(uint(seatID), updateSeatDTO.ToModel(exitingSeats)); err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, nil, "Seat updated successfully")
}

// DeleteSeat @Summary Remove a seat from the system
// @Description Removes a seat from the system
// @Tags seats
// @Accept  json
// @Produce  json
// @Param id path int true "Seat ID"
// @Success 200 {object} pkg.APIResponse "Successfully deleted seat"
// @Failure 400 {object} pkg.APIResponse "Bad Request"
// @Failure 404 {object} pkg.APIResponse "Not Found"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error"
// @Router /{busID}/seats/{id} [delete]
func (h *SeatHandler) DeleteSeat(c *gin.Context) {
	seatID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid seat ID format"))
		return
	}

	if err := h.seatService.DeleteSeat(uint(seatID)); err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, nil, "Seat deleted successfully")
}

// GetAvailableSeats @Summary Retrieve available seats
// @Description Retrieves all seats that are currently available
// @Tags seats
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.SeatResponse "Array of available seats"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error"
// @Router /{busID}/seats/availability [get]
func (h *SeatHandler) GetAvailableSeats(c *gin.Context) {
	seats, err := h.seatService.GetAvailableSeats()
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, seats, "")
}

// UpdateSeatStatus @Summary Update the status of a seat
// @Description Updates the status of a seat (booked, available, reserved)
// @Tags seats
// @Accept  json
// @Produce  json
// @Param id path int true "Seat ID"
// @Param status body dto.UpdateSeatRequest true "Seat status data"
// @Success 200 {object} pkg.APIResponse "Successfully updated seat status"
// @Failure 400 {object} pkg.APIResponse "Bad Request"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error"
// @Router /{busID}/seats/{id}/status [put]
func (h *SeatHandler) UpdateSeatStatus(c *gin.Context) {
	seatID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid seat ID format"))
		return
	}

	var statusDTO dto.UpdateSeatRequest
	if err := c.ShouldBindJSON(&statusDTO); err != nil {
		//Format the validation errors for the response.
		validationErrors := pkg.FormatValidationError(err, statusDTO)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": validationErrors,
		})
		return
	}

	if err := h.seatService.UpdateSeatStatus(uint(seatID), *statusDTO.SeatStatus); err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, nil, "Seat status updated successfully")
}

// GetSeatsByStatus @Summary Retrieve seats with a specific status
// @Description Retrieves all seats with a specific status
// @Tags seats
// @Accept  json
// @Produce  json
// @Param status path string true "Seat Status" Enums(booked, available, reserved)
// @Success 200 {array} dto.SeatResponse "Array of seats with the specified status"
// @Failure 400 {object} pkg.APIResponse "Bad Request"
// @Failure 500 {object} pkg.APIResponse "Internal Server Error"
// @Router /{busID}/seats/status/{status} [get]
func (h *SeatHandler) GetSeatsByStatus(c *gin.Context) {
	status := models.SeatStatus(c.Param("status"))
	seats, err := h.seatService.GetSeatsByStatus(status)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, seats, "")
}

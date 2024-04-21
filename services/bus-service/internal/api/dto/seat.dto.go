package dto

import "bus-service/internal/models"

// CreateSeatRequest defines the data structure for creating a new seat.
type CreateSeatRequest struct {
	BusID       uint                 `json:"bus_id" binding:"required"`
	SeatNumber  string               `json:"seat_number" binding:"required,alphanum"`
	ClassType   models.SeatClassType `json:"class_type" binding:"required,oneof=Regular Business"`
	IsAvailable bool                 `json:"is_available" binding:"omitempty"` // Optional, defaults to true if not specified.
	SeatStatus  models.SeatStatus    `json:"seat_status" binding:"required,oneof=Booked Available Reserved"`
}

// ToModel converts the CreateSeatRequest DTO to the Seat model.
func (c *CreateSeatRequest) ToModel() *models.Seat {
	return &models.Seat{
		BusID:       c.BusID,
		SeatNumber:  c.SeatNumber,
		ClassType:   c.ClassType,
		IsAvailable: c.IsAvailable,
		SeatStatus:  c.SeatStatus,
	}
}

// UpdateSeatRequest defines the data structure for updating an existing seat.
type UpdateSeatRequest struct {
	BusID       *uint                 `json:"bus_id,omitempty"` // Use pointers to differentiate between zero value and omitted field.
	SeatNumber  *string               `json:"seat_number,omitempty" binding:"omitempty,alphanum"`
	ClassType   *models.SeatClassType `json:"class_type,omitempty" binding:"omitempty,oneof=Regular Business"`
	IsAvailable *bool                 `json:"is_available,omitempty"` // Optional, can be omitted.
	SeatStatus  *models.SeatStatus    `json:"seat_status,omitempty" binding:"omitempty,oneof=Booked Available Reserved"`
}

// ToModel updates only the non-nil fields of a Seat model.
func (u *UpdateSeatRequest) ToModel(seat *models.Seat) *models.Seat {
	if u.BusID != nil {
		seat.BusID = *u.BusID
	}
	if u.SeatNumber != nil {
		seat.SeatNumber = *u.SeatNumber
	}
	if u.ClassType != nil {
		seat.ClassType = *u.ClassType
	}
	if u.IsAvailable != nil {
		seat.IsAvailable = *u.IsAvailable
	}
	if u.SeatStatus != nil {
		seat.SeatStatus = *u.SeatStatus
	}
	return seat
}

// SeatResponse is the DTO for sending seat data in HTTP responses.
type SeatResponse struct {
	ID          uint                 `json:"id"`
	BusID       uint                 `json:"bus_id"`
	SeatNumber  string               `json:"seat_number"`
	ClassType   models.SeatClassType `json:"class_type"`
	IsAvailable bool                 `json:"is_available"`
	SeatStatus  models.SeatStatus    `json:"seat_status"`
}

// FromSeatModel  transforms a Seat model into a SeatResponse DTO.
func FromSeatModel(seat models.Seat) SeatResponse {
	return SeatResponse{
		ID:          seat.ID,
		BusID:       seat.BusID,
		SeatNumber:  seat.SeatNumber,
		ClassType:   seat.ClassType,
		IsAvailable: seat.IsAvailable,
		SeatStatus:  seat.SeatStatus,
	}
}

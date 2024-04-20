package dto

import (
	"bus-service/internal/models"
	"time"

	"github.com/go-playground/validator/v10"
)

// CreateBusDTO is used when creating a new bus. It includes validation tags.
type CreateBusDTO struct {
	RouteID         uint      `json:"routeId" binding:"required"`
	BusCode         string    `json:"busCode" binding:"required,alphanumunicode,max=100"`
	Capacity        int       `json:"capacity" binding:"required,gt=0"`
	MakeModel       string    `json:"makeModel" binding:"required,max=100"`
	Year            int       `json:"year" binding:"required,gt=1900,lt=2100"` // Assuming year range from 1900 to 2100
	LicensePlate    string    `json:"licensePlate" binding:"required,alphanum,max=20"`
	Status          string    `json:"status" binding:"required,oneof='active' 'inactive' 'in_service' 'out_of_service'"`
	LastServiceDate time.Time `json:"lastServiceDate" binding:"required"`
	NextServiceDate time.Time `json:"nextServiceDate" binding:"required,gtfield=LastServiceDate"`
}

// UpdateBusDTO is used when updating a bus. It does not require all fields.
type UpdateBusDTO struct {
	RouteID         uint      `json:"routeId"`
	BusCode         string    `json:"busCode" binding:"omitempty,alphanumunicode,max=100"`
	Capacity        int       `json:"capacity" binding:"omitempty,gt=0"`
	MakeModel       string    `json:"makeModel" binding:"omitempty,max=100"`
	Year            int       `json:"year" binding:"omitempty,gt=1900,lt=2100"` // Assuming year range from 1900 to 2100
	LicensePlate    string    `json:"licensePlate" binding:"omitempty,alphanum,max=20"`
	Status          string    `json:"status" binding:"omitempty,oneof='active' 'inactive' 'in_service' 'out_of_service'"`
	LastServiceDate time.Time `json:"lastServiceDate"`
	NextServiceDate time.Time `json:"nextServiceDate" binding:"omitempty,gtfield=LastServiceDate"`
}

func ValidateCreateBusDTO(createBusDTO CreateBusDTO) error {
	validate := validator.New()
	return validate.Struct(createBusDTO)
}

func ValidateUpdateBusDTO(updateBusDTO UpdateBusDTO) error {
	validate := validator.New()
	return validate.Struct(updateBusDTO)
}

// ToModel converts CreateBusDTO to Bus model.
func (dto *CreateBusDTO) ToModel() models.Bus {
	return models.Bus{
		RouteID:         dto.RouteID,
		BusCode:         dto.BusCode,
		Capacity:        dto.Capacity,
		MakeModel:       dto.MakeModel,
		Year:            dto.Year,
		LicensePlate:    dto.LicensePlate,
		Status:          dto.Status,
		LastServiceDate: dto.LastServiceDate,
		NextServiceDate: dto.NextServiceDate,
	}
}

// ToModel converts UpdateBusDTO to Bus model.
func (dto *UpdateBusDTO) ToModel(bus *models.Bus) {
	if dto.RouteID != 0 {
		bus.RouteID = dto.RouteID
	}
	if dto.BusCode != "" {
		bus.BusCode = dto.BusCode
	}
	if dto.Capacity != 0 {
		bus.Capacity = dto.Capacity
	}
	if dto.MakeModel != "" {
		bus.MakeModel = dto.MakeModel
	}
	if dto.Year != 0 {
		bus.Year = dto.Year
	}
	if dto.LicensePlate != "" {
		bus.LicensePlate = dto.LicensePlate
	}
	if dto.Status != "" {
		bus.Status = dto.Status
	}
	if !dto.LastServiceDate.IsZero() {
		bus.LastServiceDate = dto.LastServiceDate
	}
	if !dto.NextServiceDate.IsZero() {
		bus.NextServiceDate = dto.NextServiceDate
	}
}

type BusResponse struct {
	ID              uint      `json:"id"`
	RouteID         uint      `json:"routeId"`
	BusCode         string    `json:"busCode"`
	Capacity        int       `json:"capacity"`
	MakeModel       string    `json:"makeModel"`
	Year            int       `json:"year"`
	LicensePlate    string    `json:"licensePlate"`
	Status          string    `json:"status"`
	LastServiceDate time.Time `json:"lastServiceDate"`
	NextServiceDate time.Time `json:"nextServiceDate"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// FromModel converts Bus model to BusResponseDTO.
func FromModel(bus models.Bus) BusResponse {
	return BusResponse{
		ID:              bus.ID,
		RouteID:         bus.RouteID,
		BusCode:         bus.BusCode,
		Capacity:        bus.Capacity,
		MakeModel:       bus.MakeModel,
		Year:            bus.Year,
		LicensePlate:    bus.LicensePlate,
		Status:          bus.Status,
		LastServiceDate: bus.LastServiceDate,
		NextServiceDate: bus.NextServiceDate,
		CreatedAt:       bus.CreatedAt,
		UpdatedAt:       bus.UpdatedAt,
	}
}

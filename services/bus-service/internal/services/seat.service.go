package services

import (
	"bus-service/internal/models"
	"bus-service/internal/repository"
	"errors"
)

// ISeatService defines the interface for seat-related business operations.
type ISeatService interface {
	GetSeatsByBus(busID uint) ([]models.Seat, error)
	CreateSeat(busID uint, seat *models.Seat) error
	GetSeat(seatID uint) (*models.Seat, error)
	UpdateSeat(seatID uint, seat *models.Seat) error
	DeleteSeat(seatID uint) error
	GetAvailableSeats() ([]models.Seat, error)
	UpdateSeatStatus(seatID uint, status models.SeatStatus) error
	GetSeatsByStatus(status models.SeatStatus) ([]models.Seat, error)
}

// SeatService implements the ISeatService interface for business logic related to seat management.
type SeatService struct {
	repo repository.SeatRepository
}

// NewSeatService creates a new instance of SeatService that conforms to the ISeatService interface.
func NewSeatService(repo repository.SeatRepository) ISeatService {
	return &SeatService{
		repo: repo,
	}
}

// GetSeatsByBus retrieves all seats for a given bus.
func (s *SeatService) GetSeatsByBus(busID uint) ([]models.Seat, error) {
	if busID == 0 {
		return nil, errors.New("invalid bus ID provided")
	}
	return s.repo.GetSeatsByBusID(busID)
}

// CreateSeat creates a new seat on a bus.
func (s *SeatService) CreateSeat(busID uint, seat *models.Seat) error {
	if busID == 0 || seat == nil {
		return errors.New("invalid input data provided")
	}
	if busID != seat.BusID {
		return errors.New("bus ID mismatch")
	}
	return s.repo.CreateSeat(busID, seat)
}

// GetSeat retrieves details of a specific seat.
func (s *SeatService) GetSeat(seatID uint) (*models.Seat, error) {
	if seatID == 0 {
		return nil, errors.New("invalid seat ID provided")
	}
	return s.repo.GetSeatByID(seatID)
}

// UpdateSeat updates the details of a specific seat.
func (s *SeatService) UpdateSeat(seatID uint, seat *models.Seat) error {
	if seatID == 0 || seat == nil {
		return errors.New("invalid input data provided")
	}
	return s.repo.UpdateSeat(seatID, seat)
}

// DeleteSeat removes a seat from the database.
func (s *SeatService) DeleteSeat(seatID uint) error {
	if seatID == 0 {
		return errors.New("invalid seat ID provided")
	}
	return s.repo.DeleteSeat(seatID)
}

// GetAvailableSeats retrieves all seats that are currently available for booking.
func (s *SeatService) GetAvailableSeats() ([]models.Seat, error) {
	return s.repo.GetAvailableSeats()
}

// UpdateSeatStatus changes the status of a specific seat.
func (s *SeatService) UpdateSeatStatus(seatID uint, status models.SeatStatus) error {
	if seatID == 0 {
		return errors.New("invalid seat ID provided")
	}
	return s.repo.UpdateSeatStatus(seatID, status)
}

// GetSeatsByStatus finds all seats with a given status.
func (s *SeatService) GetSeatsByStatus(status models.SeatStatus) ([]models.Seat, error) {
	return s.repo.GetSeatsByStatus(status)
}

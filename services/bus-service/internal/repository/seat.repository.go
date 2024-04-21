package repository

import (
	"bus-service/internal/models"
	"gorm.io/gorm"
	"log"
)

type SeatRepository interface {
	GetSeatsByBusID(busID uint) ([]models.Seat, error)
	CreateSeat(busID uint, seat *models.Seat) error
	GetSeatByID(seatID uint) (*models.Seat, error)
	UpdateSeat(seatID uint, seat *models.Seat) error
	DeleteSeat(seatID uint) error
	GetAvailableSeats() ([]models.Seat, error)
	UpdateSeatStatus(seatID uint, status models.SeatStatus) error
	GetSeatsByStatus(status models.SeatStatus) ([]models.Seat, error)
}
type GormSeatRepository struct {
	db *gorm.DB
}

func NewSeatRepository(db *gorm.DB) *GormSeatRepository {
	return &GormSeatRepository{db: db}
}

func (r *GormSeatRepository) GetSeatsByBusID(busID uint) ([]models.Seat, error) {
	var seats []models.Seat
	if err := r.db.Where("bus_id = ?", busID).Find(&seats).Error; err != nil {
		log.Print(`er`, err)
		return nil, err
	}
	return seats, nil
}

func (r *GormSeatRepository) CreateSeat(busID uint, seat *models.Seat) error {
	seat.BusID = busID
	return r.db.Create(seat).Error
}

func (r *GormSeatRepository) GetSeatByID(seatID uint) (*models.Seat, error) {
	var seat models.Seat
	if err := r.db.First(&seat, seatID).Error; err != nil {
		return nil, err
	}
	return &seat, nil
}

func (r *GormSeatRepository) UpdateSeat(seatID uint, seat *models.Seat) error {
	return r.db.Model(&models.Seat{}).Where("id = ?", seatID).Updates(seat).Error
}

func (r *GormSeatRepository) DeleteSeat(seatID uint) error {
	return r.db.Delete(&models.Seat{}, seatID).Error
}

func (r *GormSeatRepository) GetAvailableSeats() ([]models.Seat, error) {
	var seats []models.Seat
	if err := r.db.Where("is_available = ?", true).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}

func (r *GormSeatRepository) UpdateSeatStatus(seatID uint, status models.SeatStatus) error {
	return r.db.Model(&models.Seat{}).Where("id = ?", seatID).Update("seat_status", status).Error
}

func (r *GormSeatRepository) GetSeatsByStatus(status models.SeatStatus) ([]models.Seat, error) {
	var seats []models.Seat
	if err := r.db.Where("seat_status = ?", status).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}

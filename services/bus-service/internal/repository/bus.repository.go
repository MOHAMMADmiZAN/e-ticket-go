package repository

import (
	"bus-service/internal/models"
	_ "errors"
	_ "fmt"
	"gorm.io/gorm"
	"time"
)

type BusRepository struct {
	DB *gorm.DB
}

// NewBusRepository creates a new instance of BusRepository.
func NewBusRepository(db *gorm.DB) *BusRepository {
	return &BusRepository{DB: db}
}

// GetAllBuses retrieves all buses from the database.
func (repo *BusRepository) GetAllBuses() ([]models.Bus, error) {
	var buses []models.Bus
	result := repo.DB.Find(&buses)
	return buses, result.Error
}

// GetBusByID retrieves a bus by its ID.
func (repo *BusRepository) GetBusByID(id uint) (*models.Bus, error) {
	var bus models.Bus
	result := repo.DB.First(&bus, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bus, nil
}

// CreateBus adds a new bus to the database.
func (repo *BusRepository) CreateBus(bus models.Bus) (*models.Bus, error) {
	result := repo.DB.Create(&bus)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bus, nil
}

// UpdateBus updates an existing bus.
func (repo *BusRepository) UpdateBus(bus models.Bus) (*models.Bus, error) {
	// Find the existing bus record. We don't want to overwrite fields with zero values.
	var existingBus models.Bus
	if err := repo.DB.First(&existingBus, bus.ID).Error; err != nil {
		return nil, err // Bus not found or other error.
	}

	// Map the changes from the provided bus object onto the existing bus.
	// Omit any fields from bus that should not be updated.
	if err := repo.DB.Model(&existingBus).Updates(bus).Error; err != nil {
		return nil, err
	}

	return &existingBus, nil
}

// DeleteBus removes a bus from the database.
func (repo *BusRepository) DeleteBus(id uint) error {
	result := repo.DB.Delete(&models.Bus{}, id)
	return result.Error
}

// GetBusesByStatus retrieves all buses with a specific status.
func (repo *BusRepository) GetBusesByStatus(status string) ([]models.Bus, error) {
	var buses []models.Bus
	result := repo.DB.Where("status = ?", status).Find(&buses)
	return buses, result.Error
}

// UpdateBusServiceDates updates the last and next service dates for a bus.
func (repo *BusRepository) UpdateBusServiceDates(id uint, lastServiceDate, nextServiceDate time.Time) error {
	result := repo.DB.Model(&models.Bus{}).Where("id = ?", id).Updates(map[string]interface{}{
		"last_service_date": lastServiceDate,
		"next_service_date": nextServiceDate,
	})
	return result.Error
}

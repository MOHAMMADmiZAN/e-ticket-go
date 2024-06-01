package repository

import (
	"bus-service/internal/models"
	"errors"
	_ "errors"
	_ "fmt"
	"time"

	"gorm.io/gorm"
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

// UpdateBusServiceDates updates the last and next service dates for a bus with given ID after performing necessary validations.
func (repo *BusRepository) UpdateBusServiceDates(id uint, lastServiceDate, nextServiceDate time.Time) error {
	if time.Now().Before(lastServiceDate) {
		return errors.New("last service date cannot be in the future")
	}

	if !nextServiceDate.After(lastServiceDate) {
		return errors.New("next service date must be after the last service date")
	}

	// Retrieve the existing bus to check if it exists before updating
	var existingBus models.Bus
	if err := repo.DB.Where("id = ?", id).First(&existingBus).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no bus found with the given ID")
		}
		return err
	}

	// Update the bus with new service dates
	result := repo.DB.Model(&existingBus).Updates(models.Bus{
		LastServiceDate: lastServiceDate,
		NextServiceDate: nextServiceDate,
	})

	// Check for errors in the update operation
	if result.Error != nil {
		return result.Error
	}

	// Check if the update operation actually modified any rows
	if result.RowsAffected == 0 {
		return errors.New("no updates performed; the bus data might already be up to date")
	}

	return nil
}

// GetBusByCodeOrLicensePlate retrieves a Bus by its BusCode or License Plate.
func (repo *BusRepository) GetBusByCodeOrLicensePlate(busCode, licensePlate string) (*models.Bus, error) {
	var bus models.Bus
	// Build a query using an OR condition to check either bus_code or license_plate
	result := repo.DB.Where("bus_code = ? OR license_plate = ?", busCode, licensePlate).First(&bus)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bus, nil
}

// GetBusesByRouteID retrieves all buses assigned to a specific route.
func (repo *BusRepository) GetBusesByRouteID(routeID uint) ([]models.Bus, error) {
	var buses []models.Bus
	result := repo.DB.Where("route_id = ?", routeID).Find(&buses)
	return buses, result.Error
}

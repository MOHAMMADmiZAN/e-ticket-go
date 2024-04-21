package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Bus struct {
	gorm.Model
	RouteID         uint      `gorm:"not null;index"` // Assume validation through RouteService
	BusCode         string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	Capacity        int       `gorm:"not null;"`
	MakeModel       string    `gorm:"type:varchar(100);not null"`
	Year            int       `gorm:"not null;"`
	LicensePlate    string    `gorm:"type:varchar(20);not null;unique"`
	Status          string    `gorm:"type:varchar(50);not null;default:'active'"`
	LastServiceDate time.Time `gorm:"not null"`
	NextServiceDate time.Time `gorm:"not null"`
	Seats           []Seat    `gorm:"foreignKey:BusID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-many relationship with Seats
}

// TableName specifies the table name for Bus
func (Bus) TableName() string {
	return "buses"
}

// BeforeCreate contains all the custom validation logic before creating a Bus record
func (b *Bus) BeforeCreate(tx *gorm.DB) error {
	// Assume we have a RouteServiceClient for validating the RouteID
	if err := validateRouteID(b.RouteID); err != nil {
		return err
	}

	if b.BusCode == "" {
		return errors.New("bus code must not be empty")
	}

	if len(b.BusCode) > 100 {
		return errors.New("bus code must not exceed 100 characters")
	}

	if b.Capacity <= 0 {
		return errors.New("capacity must be positive")
	}

	if b.MakeModel == "" {
		return errors.New("make and models must not be empty")
	}

	if len(b.MakeModel) > 100 {
		return errors.New("make and models must not exceed 100 characters")
	}

	if b.Year < 1900 || b.Year > time.Now().Year() {
		return fmt.Errorf("year must be between 1900 and the current year, got %d", b.Year)
	}

	if b.LicensePlate == "" {
		return errors.New("license plate must not be empty")
	}

	if len(b.LicensePlate) > 20 {
		return errors.New("license plate must not exceed 20 characters")
	}

	if !isValidBusStatus(b.Status) {
		return fmt.Errorf("status must be one of the predefined values, got '%s'", b.Status)
	}

	if b.LastServiceDate.IsZero() || b.LastServiceDate.After(time.Now()) {
		return errors.New("last services date must be in the past and not empty")
	}

	if b.NextServiceDate.IsZero() || b.NextServiceDate.Before(time.Now()) {
		return errors.New("next services date must be in the future and not empty")
	}

	return nil
}

// BeforeUpdate contains all the custom validation logic before updating a Bus record
func (b *Bus) BeforeUpdate(tx *gorm.DB) error {
	// Reuse the same validation logic for update
	return b.BeforeCreate(tx)
}

func validateRouteID(routeID uint) error {

	// Mock validation
	if routeID == 0 {
		return errors.New("route ID must be valid and non-zero")
	}

	return nil
}

func isValidBusStatus(status string) bool {
	switch status {
	case "active", "maintenance", "decommissioned":
		return true
	default:
		return false
	}
}

package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"regexp"
)

// SeatClassType and SeatStatus define allowed enum values for the ClassType and SeatStatus fields respectively.
type SeatClassType string
type SeatStatus string

const (
	// ClassRegular Definitions for class types
	ClassRegular  SeatClassType = "Regular"
	ClassBusiness SeatClassType = "Business"

	// StatusBooked Definitions for seat statuses
	StatusBooked    SeatStatus = "Booked"
	StatusAvailable SeatStatus = "Available"
	StatusReserved  SeatStatus = "Reserved"
)

// Seat represents a seat on a bus, with ORM functionalities managed by GORM.
type Seat struct {
	gorm.Model
	ID          uint          `gorm:"primaryKey;autoIncrement"`
	BusID       uint          `gorm:"index;not null;constraint:OnDelete:CASCADE"`           // Foreign key referencing the buses table.
	SeatNumber  string        `gorm:"size:255;not null;uniqueIndex:idx_seat_number_bus_id"` // Alphanumeric identifier for the seat, unique within the bus
	ClassType   SeatClassType `gorm:"type:varchar(100);not null"`                           // Class type of the seat.
	IsAvailable bool          `gorm:"default:true"`                                         // Indicates if the seat is available for booking.
	SeatStatus  SeatStatus    `gorm:"type:varchar(100);not null"`                           // Current status of the seat.
}

// TableName specifies the table name for GORM to use, overriding the default.
func (Seat) TableName() string {
	return "seats"
}

// validateSeatNumber checks if the seat number is alphanumeric and conforms to expected standards.
func validateSeatNumber(number string) error {
	if number == "" {
		return errors.New("seat number cannot be empty")
	}
	matched, err := regexp.MatchString(`^[a-zA-Z0-9]+$`, number)
	if err != nil {
		return fmt.Errorf("error validating seat number: %v", err)
	}
	if !matched {
		return errors.New("seat number must be alphanumeric")
	}
	return nil
}

// BeforeSave is a GORM hook executed before saving a Seat record (both insert and update).
func (s *Seat) BeforeSave(tx *gorm.DB) error {
	if err := validateSeatNumber(s.SeatNumber); err != nil {
		return err
	}
	if s.ClassType != ClassRegular && s.ClassType != ClassBusiness {
		return fmt.Errorf("invalid class type: %s", s.ClassType)
	}
	if s.SeatStatus != StatusBooked && s.SeatStatus != StatusAvailable && s.SeatStatus != StatusReserved {
		return fmt.Errorf("invalid seat status: %s", s.SeatStatus)
	}
	return nil
}

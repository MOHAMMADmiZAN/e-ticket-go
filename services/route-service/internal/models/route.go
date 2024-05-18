package models

import (
	"time"

	"gorm.io/gorm"
)

// Route represents the transportation route.
type Route struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey autoIncrement"`
	Name          string    `gorm:"index;not null;type:varchar(100)"` // Explicitly set the column type.
	StartTime     time.Time `gorm:"not null"`
	Duration      int       `gorm:"not null"`
	StartLocation string    `json:"start_location" gorm:"not null"`
	EndLocation   string    `json:"end_location" gorm:"not null"`
	Stops         []Stop    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (r *Route) AfterDelete(tx *gorm.DB) error {
	// Delete related Stops
	if err := tx.Where("route_id = ?", r.ID).Delete(&Stop{}).Error; err != nil {
		return err
	}

	return nil
}
func (Route) TableName() string {
	return "routes"
}

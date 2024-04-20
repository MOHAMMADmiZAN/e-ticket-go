package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

// Route represents the transportation route.
type Route struct {
	ID            uint           `gorm:"primaryKey"`
	Name          string         `gorm:"index;not null;type:varchar(100)"` // Explicitly set the column type.
	StartTime     time.Time      `gorm:"not null"`
	Duration      int            `gorm:"not null"`
	StartLocation string         `json:"start_location" gorm:"not null"`
	EndLocation   string         `json:"end_location" gorm:"not null"`
	Stops         []Stop         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Schedules     []Schedule     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time      `gorm:"autoCreateTime default:current_timestamp"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime default:current_timestamp"`
	DeletedAt     gorm.DeletedAt ` gorm:"index"`
}

func (r *Route) AfterDelete(tx *gorm.DB) error {
	// Delete related Stops
	if err := tx.Where("route_id = ?", r.ID).Delete(&Stop{}).Error; err != nil {
		return err
	}

	// Delete related Schedules
	if err := tx.Where("route_id = ?", r.ID).Delete(&Schedule{}).Error; err != nil {
		return err
	}

	return nil
}
func (r *Route) AfterFind(tx *gorm.DB) (err error) {
	// Custom logic to execute after a route is found
	log.Printf("Route with ID %d has been found", r.ID)
	// You can perform additional operations or transformations here
	return nil
}

func (Route) TableName() string {
	return "routes"
}

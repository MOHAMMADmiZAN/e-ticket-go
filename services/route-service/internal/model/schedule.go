package model

import (
	"gorm.io/gorm"
	"time"
)

// Schedule represents the timetable for a route at a specific stop.
type Schedule struct {
	ID            uint      `gorm:"primaryKey"`
	RouteID       uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StopID        uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ArrivalTime   time.Time `gorm:"not null"`
	DepartureTime time.Time `gorm:"not null"`
	Route         Route     `gorm:"foreignkey:RouteID"`
	Stop          Stop      `gorm:"foreignkey:StopID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

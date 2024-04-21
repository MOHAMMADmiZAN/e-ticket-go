package models

import (
	"gorm.io/gorm"
	"time"
)

// Schedule represents the timetable for a route at a specific stop.
type Schedule struct {
	gorm.Model
	RouteID       uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StopID        uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ArrivalTime   time.Time `gorm:"not null"`
	DepartureTime time.Time `gorm:"not null"`
	Route         Route     `gorm:"foreignkey:RouteID"`
	Stop          Stop      `gorm:"foreignkey:StopID"`
}

func (Schedule) TableName() string {
	return "schedules"

}

package models

import (
	"gorm.io/gorm"
	"time"
)

// Schedule represents the timetable for a route at a specific stop.
type Schedule struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey autoIncrement"`
	StopID        uint      `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ArrivalTime   time.Time `gorm:"not null"`
	DepartureTime time.Time `gorm:"not null"`
	Stop          Stop      `gorm:"foreignkey:StopID"`
}

func (Schedule) TableName() string {
	return "schedules"

}

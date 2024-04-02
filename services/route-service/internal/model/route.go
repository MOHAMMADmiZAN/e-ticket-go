package model

import (
	"gorm.io/gorm"
	"time"
)

// Route represents the transportation route.
type Route struct {
	ID            uint       `gorm:"primaryKey"`
	Name          string     `gorm:"index;not null;type:varchar(100)"` // Explicitly set the column type.
	StartTime     time.Time  `gorm:"not null"`
	Duration      int        `gorm:"not null"`
	StartLocation string     `json:"start_location" gorm:"not null"`
	EndLocation   string     `json:"end_location" gorm:"not null"`
	Stops         []Stop     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Schedules     []Schedule `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt ` gorm:"index"`
}

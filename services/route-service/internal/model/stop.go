package model

import (
	"gorm.io/gorm"
	"time"
)

// Stop represents a bus stop along a route.
type Stop struct {
	ID        uint   `gorm:"primaryKey"`
	RouteID   uint   `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name      string `gorm:"not null;type:varchar(100)"`
	Sequence  int    `gorm:"not null"`
	Route     Route  `gorm:"foreignkey:RouteID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

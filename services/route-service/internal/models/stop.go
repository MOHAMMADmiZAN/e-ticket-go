package models

import (
	"gorm.io/gorm"
)

// Stop represents a bus stop along a route.
type Stop struct {
	gorm.Model
	RouteID  uint   `gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name     string `gorm:"not null;type:varchar(100)"`
	Sequence int    `gorm:"not null"`
	Route    Route  `gorm:"foreignkey:RouteID"`
}

func (Stop) TableName() string {
	return "stops"
}

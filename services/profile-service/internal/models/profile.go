package models

import (
	"gorm.io/gorm"
	"time"
)

type UserProfile struct {
	gorm.Model
	UserID            uint      `gorm:"index;not null" json:"userID"` // Foreign key for User model
	FirstName         string    `gorm:"size:100;not null" json:"firstName" validate:"required,alpha"`
	LastName          string    `gorm:"size:100;not null" json:"lastName" validate:"required,alpha"`
	DateOfBirth       time.Time `json:"dateOfBirth" validate:"lte"`
	ProfilePictureURL string    `gorm:"size:255" json:"profilePictureURL" validate:"omitempty,url"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

// TableName overrides the table name used by UserProfile to `user_profiles`
func (UserProfile) TableName() string {
	return "user_profiles"
}

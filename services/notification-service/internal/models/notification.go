package models

import (
	"gorm.io/gorm"
	"time"
)

// Notification represents the notifications sent to users.
type Notification struct {
	gorm.Model
	UserID       uint      `gorm:"not null;index" json:"userID"` // Foreign key for User model
	Type         string    `gorm:"size:100;not null" json:"type"`
	Status       string    `gorm:"size:100;not null" json:"status"`
	Channel      string    `gorm:"size:100;not null" json:"channel"`
	Content      string    `gorm:"type:text;not null" json:"content"`
	SendDate     time.Time `json:"sendDate"`
	Acknowledged bool      `json:"acknowledged"`
}

// TableName specifies the table name for GORM.
func (Notification) TableName() string {
	return "notifications"
}

// UserNotificationPreferences represents a user's notification preferences.
type UserNotificationPreferences struct {
	gorm.Model
	UserID       uint   `gorm:"not null;index" json:"userID"` // Foreign key for User model
	PrefersEmail bool   `json:"prefersEmail"`
	PrefersSMS   bool   `json:"prefersSMS"`
	Email        string `gorm:"size:255" json:"email"`
	PhoneNumber  string `gorm:"size:255" json:"phoneNumber"`
}

// TableName specifies the table name for GORM.
func (UserNotificationPreferences) TableName() string {
	return "user_notification_preferences"
}

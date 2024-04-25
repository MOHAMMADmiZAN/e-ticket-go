package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log" // Make sure to import the log package
	"notification-service/internal/models"
)

// UserNotification defines the interface for user notification preference storage operations.
type UserNotification interface {
	CreatePreferences(prefs models.UserNotificationPreferences) error
	GetPreferences(userID uint) (models.UserNotificationPreferences, error)
	UpdatePreferences(prefs models.UserNotificationPreferences) error
	DeletePreferences(userID uint) error
}

// UserNotificationRepositoryImpl  implements UserRepository with a GORM backend.
type UserNotificationRepositoryImpl struct {
	db *gorm.DB
}

// NewUserNotificationRepository initializes a new UserRepository with the given GORM DB connection.
func NewUserNotificationRepository(db *gorm.DB) UserNotification {
	return &UserNotificationRepositoryImpl{db: db}
}

// CreatePreferences saves new user notification preferences in the database.
func (repo *UserNotificationRepositoryImpl) CreatePreferences(prefs models.UserNotificationPreferences) error {
	result := repo.db.Create(&prefs)
	if result.Error != nil {
		log.Printf("Error creating user notification preferences: %v", result.Error)
		return fmt.Errorf("failed to create user notification preferences: %w", result.Error)
	}
	return nil
}

// GetPreferences retrieves the notification preferences for a given user ID.
func (repo *UserNotificationRepositoryImpl) GetPreferences(userID uint) (models.UserNotificationPreferences, error) {
	var prefs models.UserNotificationPreferences
	result := repo.db.Where("user_id = ?", userID).First(&prefs)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("User preferences not found for user ID %d", userID)
			return prefs, gorm.ErrRecordNotFound
		}
		log.Printf("Error retrieving user notification preferences: %v", result.Error)
		return prefs, fmt.Errorf("failed to retrieve user notification preferences: %w", result.Error)
	}
	return prefs, nil
}

// UpdatePreferences updates existing notification preferences in the database.
func (repo *UserNotificationRepositoryImpl) UpdatePreferences(prefs models.UserNotificationPreferences) error {
	result := repo.db.Save(&prefs)
	if result.Error != nil {
		log.Printf("Error updating user notification preferences: %v", result.Error)
		return fmt.Errorf("failed to update user notification preferences: %w", result.Error)
	}
	return nil
}

// DeletePreferences removes the notification preferences for a given user ID.
func (repo *UserNotificationRepositoryImpl) DeletePreferences(userID uint) error {
	result := repo.db.Where("user_id = ?", userID).Delete(&models.UserNotificationPreferences{})
	if result.Error != nil {
		log.Printf("Error deleting user notification preferences for user ID %d: %v", userID, result.Error)
		return fmt.Errorf("failed to delete user notification preferences: %w", result.Error)
	}
	return nil
}

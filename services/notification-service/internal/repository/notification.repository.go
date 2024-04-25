package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"notification-service/internal/models"
)

// INotificationRepository defines the interface for the repository handling notifications.
type INotificationRepository interface {
	Create(notification *models.Notification) error
	Delete(notificationID uint) error
	GetByID(notificationID uint) (*models.Notification, error)
	List(userID uint, offset, limit int) ([]*models.Notification, error)
}

// NotificationRepository implements the operations defined in INotificationRepository.
type NotificationRepository struct {
	db *gorm.DB
}

// NewNotificationRepository creates a new instance of a notification repository.
func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db}
}

// Create inserts a new notification into the database.
func (r *NotificationRepository) Create(notification *models.Notification) error {
	if err := r.db.Create(notification).Error; err != nil {
		return fmt.Errorf("creating notification failed: %w", err)
	}
	return nil
}

// Delete removes a notification from the database by ID.
func (r *NotificationRepository) Delete(notificationID uint) error {
	result := r.db.Delete(&models.Notification{}, notificationID)
	if result.Error != nil {
		return fmt.Errorf("deleting notification failed: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("no notification found to delete")
	}
	return nil
}

// GetByID fetches a single notification by ID from the database.
func (r *NotificationRepository) GetByID(notificationID uint) (*models.Notification, error) {
	var notification models.Notification
	result := r.db.First(&notification, notificationID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("notification with ID %d not found", notificationID)
		}
		return nil, fmt.Errorf("retrieving notification failed: %w", result.Error)
	}
	return &notification, nil
}

// List retrieves a list of notifications, optionally filtered by a user ID.
func (r *NotificationRepository) List(userID uint, offset, limit int) ([]*models.Notification, error) {
	var notifications []*models.Notification
	query := r.db.Offset(offset).Limit(limit)
	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	}
	if err := query.Find(&notifications).Error; err != nil {
		return nil, fmt.Errorf("listing notifications failed: %w", err)
	}
	return notifications, nil
}

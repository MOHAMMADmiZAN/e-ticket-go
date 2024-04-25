package services

import (
	"fmt"
	"notification-service/internal/api/dto"
	"notification-service/internal/repository"
)

// INotificationService defines the interface for notification service operations.
type INotificationService interface {
	CreateNotification(req dto.CreateNotificationRequest) (*dto.NotificationResponse, error)
	DeleteNotification(id uint) error
	GetNotification(id uint) (*dto.NotificationResponse, error)
	ListNotifications(userID uint, offset, limit int) ([]*dto.NotificationResponse, error)
}

// NotificationService implements the INotificationService interface.
type NotificationService struct {
	repo repository.INotificationRepository
}

// NewNotificationService returns a new instance of a notification service.
func NewNotificationService(repo repository.INotificationRepository) *NotificationService {
	return &NotificationService{repo}
}

// CreateNotification handles the creation of a new notification.
func (s *NotificationService) CreateNotification(req dto.CreateNotificationRequest) (*dto.NotificationResponse, error) {
	notification := req.ToModel()
	if err := s.repo.Create(&notification); err != nil {
		return nil, fmt.Errorf("failed to create notification: %w", err)
	}
	return dto.FromModel(notification), nil
}

// DeleteNotification handles the deletion of a notification.
func (s *NotificationService) DeleteNotification(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete notification: %w", err)
	}
	return nil
}

// GetNotification retrieves a notification by its ID.
func (s *NotificationService) GetNotification(id uint) (*dto.NotificationResponse, error) {
	notification, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get notification: %w", err)
	}
	return dto.FromModel(*notification), nil
}

// ListNotifications retrieves a list of notifications for a given user.
func (s *NotificationService) ListNotifications(userID uint, offset, limit int) ([]*dto.NotificationResponse, error) {
	notifications, err := s.repo.List(userID, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to list notifications: %w", err)
	}

	responses := make([]*dto.NotificationResponse, len(notifications))
	for i, n := range notifications {
		responses[i] = dto.FromModel(*n)
	}
	return responses, nil
}

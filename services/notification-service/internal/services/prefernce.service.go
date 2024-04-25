package services

import (
	"fmt"
	"notification-service/internal/api/dto"
	"notification-service/internal/repository"
)

// UserNotificationService defines the interface for user notification preference service operations.
type UserNotificationService interface {
	CreateUserPreferences(userID uint, req dto.UserNotificationPreferencesRequest) error
	UpdateUserPreferences(userID uint, update dto.UserNotificationPreferencesUpdate) error
	GetUserPreferences(userID uint) (dto.UserNotificationPreferencesResponse, error)
	DeleteUserPreferences(userID uint) error
}

type userNotificationServiceImpl struct {
	userRepo repository.UserNotification
}

// NewUserNotificationService returns a new instance of a user notification service.
func NewUserNotificationService(userRepo repository.UserNotification) UserNotificationService {
	return &userNotificationServiceImpl{userRepo: userRepo}
}

// CreateUserPreferences creates new user notification preferences with validation and business logic.
func (s *userNotificationServiceImpl) CreateUserPreferences(userID uint, req dto.UserNotificationPreferencesRequest) error {
	if req.Email == "" && req.PrefersEmail {
		return fmt.Errorf("email address is required when email notifications are preferred")
	}

	prefs := req.ToPreferencesModel(userID)
	if err := s.userRepo.CreatePreferences(prefs); err != nil {
		return fmt.Errorf("error creating user preferences: %w", err)
	}
	return nil
}

// UpdateUserPreferences updates user's notification preferences with business logic checks.
func (s *userNotificationServiceImpl) UpdateUserPreferences(userID uint, update dto.UserNotificationPreferencesUpdate) error {
	currentPrefs, err := s.userRepo.GetPreferences(userID)
	if err != nil {
		return fmt.Errorf("error retrieving current preferences: %w", err)
	}

	// Additional business logic before applying updates
	if update.Email != nil && *update.Email == "" && *update.PrefersEmail {
		return fmt.Errorf("email address is required when email notifications are preferred")
	}

	update.ToPreferencesUpdateModel(&currentPrefs)
	if err := s.userRepo.UpdatePreferences(currentPrefs); err != nil {
		return fmt.Errorf("error updating user preferences: %w", err)
	}
	return nil
}

// GetUserPreferences retrieves the user notification preferences and ensures valid response structuring.
func (s *userNotificationServiceImpl) GetUserPreferences(userID uint) (dto.UserNotificationPreferencesResponse, error) {
	prefs, err := s.userRepo.GetPreferences(userID)
	if err != nil {
		return dto.UserNotificationPreferencesResponse{}, fmt.Errorf("error fetching preferences: %w", err)
	}
	return dto.FromPreferencesModel(prefs), nil
}

// DeleteUserPreferences deletes a user's notification preferences with confirmation of deletion.
func (s *userNotificationServiceImpl) DeleteUserPreferences(userID uint) error {
	if err := s.userRepo.DeletePreferences(userID); err != nil {
		return fmt.Errorf("error deleting user preferences: %w", err)
	}
	return nil
}

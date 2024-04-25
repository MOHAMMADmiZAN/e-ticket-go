package dto

import (
	"notification-service/internal/models"
	"time"
)

// UserNotificationPreferencesRequest defines the structure for updating user notification preferences.
type UserNotificationPreferencesRequest struct {
	PrefersEmail bool   `json:"prefersEmail"`
	PrefersSMS   bool   `json:"prefersSMS"`
	Email        string `json:"email" validate:"omitempty,email"`
	PhoneNumber  string `json:"phoneNumber" validate:"omitempty"`
}

// ToPreferencesModel converts UserNotificationPreferencesRequest to UserNotificationPreferences model.
// It accepts a userID to be set, ensuring the model is associated with the correct user.
func (req *UserNotificationPreferencesRequest) ToPreferencesModel(userID uint) models.UserNotificationPreferences {
	return models.UserNotificationPreferences{
		UserID:       userID,
		PrefersEmail: req.PrefersEmail,
		PrefersSMS:   req.PrefersSMS,
		Email:        req.Email,
		PhoneNumber:  req.PhoneNumber,
	}
}

// UserNotificationPreferencesUpdate defines the structure for partially updating user notification preferences.
type UserNotificationPreferencesUpdate struct {
	PrefersEmail *bool   `json:"prefersEmail,omitempty"`
	PrefersSMS   *bool   `json:"prefersSMS,omitempty"`
	Email        *string `json:"email,omitempty" validate:"omitempty,email"`
	PhoneNumber  *string `json:"phoneNumber,omitempty" validate:"omitempty"`
}

// ToPreferencesUpdateModel applies update fields to the existing UserNotificationPreferences model.
func (update *UserNotificationPreferencesUpdate) ToPreferencesUpdateModel(model *models.UserNotificationPreferences) {
	if update.PrefersEmail != nil {
		model.PrefersEmail = *update.PrefersEmail
	}
	if update.PrefersSMS != nil {
		model.PrefersSMS = *update.PrefersSMS
	}
	if update.Email != nil {
		model.Email = *update.Email
	}
	if update.PhoneNumber != nil {
		model.PhoneNumber = *update.PhoneNumber
	}
	model.UpdatedAt = time.Now()
}

// UserNotificationPreferencesResponse defines the structure for a response that includes user notification preferences.
type UserNotificationPreferencesResponse struct {
	UserID       uint   `json:"userID"`
	PrefersEmail bool   `json:"prefersEmail"`
	PrefersSMS   bool   `json:"prefersSMS"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
}

// FromPreferencesModel converts UserNotificationPreferences model to UserNotificationPreferencesResponse.
func FromPreferencesModel(model models.UserNotificationPreferences) UserNotificationPreferencesResponse {
	return UserNotificationPreferencesResponse{
		UserID:       model.UserID,
		PrefersEmail: model.PrefersEmail,
		PrefersSMS:   model.PrefersSMS,
		Email:        model.Email,
		PhoneNumber:  model.PhoneNumber,
	}
}

package dto

import (
	"profile-service/internal/models"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// UserProfileRequest is used when creating a new user profile.
type UserProfileRequest struct {
	UserID            uint      `json:"userID" validate:"required"`
	FirstName         string    `json:"firstName" validate:"required,alpha"`
	LastName          string    `json:"lastName" validate:"required,alpha"`
	DateOfBirth       time.Time `json:"dateOfBirth" validate:"required,lte"`
	ProfilePictureURL string    `json:"profilePictureURL" validate:"omitempty,url"`
}

// ToUserProfileModel transforms UserProfileRequest to UserProfile model.
func (ur UserProfileRequest) ToUserProfileModel() (models.UserProfile, error) {
	if err := validate.Struct(ur); err != nil {
		return models.UserProfile{}, err
	}
	return models.UserProfile{
		UserID:            ur.UserID,
		FirstName:         ur.FirstName,
		LastName:          ur.LastName,
		DateOfBirth:       ur.DateOfBirth,
		ProfilePictureURL: ur.ProfilePictureURL,
	}, nil
}

// UserProfileUpdate is used when updating an existing user profile.
type UserProfileUpdate struct {
	FirstName         *string    `json:"firstName" validate:"omitempty,alpha"`
	LastName          *string    `json:"lastName" validate:"omitempty,alpha"`
	DateOfBirth       *time.Time `json:"dateOfBirth" validate:"omitempty,lte"`
	ProfilePictureURL *string    `json:"profilePictureURL" validate:"omitempty,url"`
}

// ToModel transforms UserProfileUpdate to UserProfile model for updating.
func (uu UserProfileUpdate) ToModel(id uint) (models.UserProfile, error) {
	if err := validate.Struct(uu); err != nil {
		return models.UserProfile{}, err
	}
	userProfile := models.UserProfile{
		UserID: id,
	}
	if uu.FirstName != nil {
		userProfile.FirstName = *uu.FirstName
	}
	if uu.LastName != nil {
		userProfile.LastName = *uu.LastName
	}
	if uu.DateOfBirth != nil {
		userProfile.DateOfBirth = *uu.DateOfBirth
	}
	if uu.ProfilePictureURL != nil {
		userProfile.ProfilePictureURL = *uu.ProfilePictureURL
	}
	userProfile.UpdatedAt = time.Now()
	return userProfile, nil
}

// UserProfileResponse is used to provide a user profile data to the client.
type UserProfileResponse struct {
	UserID            uint      `json:"userID"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	DateOfBirth       time.Time `json:"dateOfBirth"`
	ProfilePictureURL string    `json:"profilePictureURL"`
}

// FromUserProfileModel transforms UserProfile model to UserProfileResponse.
func FromUserProfileModel(u models.UserProfile) UserProfileResponse {
	return UserProfileResponse{
		UserID:            u.UserID,
		FirstName:         u.FirstName,
		LastName:          u.LastName,
		DateOfBirth:       u.DateOfBirth,
		ProfilePictureURL: u.ProfilePictureURL,
	}
}

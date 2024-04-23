package services

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"profile-service/internal/api/dto"
	"profile-service/internal/repository"
)

// Custom error types for specific error handling by the service consumers.
type (
	ErrUserNotFound struct {
		UserID uint
	}
	ErrUserProfileAlreadyExists struct {
		UserID uint
	}
)

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("user with ID %v not found", e.UserID)
}

func (e *ErrUserProfileAlreadyExists) Error() string {
	return fmt.Sprintf("user profile with ID %v already exists", e.UserID)
}

// UserProfileService is responsible for handling user profile-related business logic.
type UserProfileService struct {
	repo     repository.IUserProfileRepository
	validate *validator.Validate
}
type IUserProfileService interface {
	CreateUserProfile(request dto.UserProfileRequest) (*dto.UserProfileResponse, error)
	GetUserProfile(userID uint) (*dto.UserProfileResponse, error)
	UpdateUserProfile(userID uint, request dto.UserProfileUpdate) (*dto.UserProfileResponse, error)
	DeleteUserProfile(userID uint) error
}

// NewUserProfileService creates a new instance of user profile service.
func NewUserProfileService(repo repository.IUserProfileRepository) IUserProfileService {
	return &UserProfileService{
		repo:     repo,
		validate: validator.New(),
	}
}

// CreateUserProfile manages the creation of a new user profile.
func (s *UserProfileService) CreateUserProfile(request dto.UserProfileRequest) (*dto.UserProfileResponse, error) {
	// Business validation: ensure user does not already have a profile.
	existingProfile, _ := s.repo.GetByUserID(request.UserID)
	if existingProfile != nil {
		return nil, &ErrUserProfileAlreadyExists{UserID: request.UserID}
	}

	// Data validation
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	profileModel, err := request.ToUserProfileModel()
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(&profileModel); err != nil {
		return nil, err
	}

	response := dto.FromUserProfileModel(profileModel)
	return &response, nil
}

// GetUserProfile retrieves the user profile by ID with proper error handling.
func (s *UserProfileService) GetUserProfile(userID uint) (*dto.UserProfileResponse, error) {
	profile, err := s.repo.GetByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &ErrUserNotFound{UserID: userID}
		}
		return nil, err
	}

	response := dto.FromUserProfileModel(*profile)
	return &response, nil
}

// UpdateUserProfile manages user profile updates with comprehensive validation.
func (s *UserProfileService) UpdateUserProfile(userID uint, request dto.UserProfileUpdate) (*dto.UserProfileResponse, error) {
	// Business validation: ensure the profile exists.
	_, err := s.repo.GetByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &ErrUserNotFound{UserID: userID}
		}
		return nil, err
	}

	// Data validation
	if err := s.validate.Struct(request); err != nil {
		return nil, err
	}

	profileModel, err := request.ToModel(userID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Update(userID, profileModel); err != nil {
		return nil, err
	}

	// Fetch the updated profile to include in the response.
	updatedProfile, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := dto.FromUserProfileModel(*updatedProfile)
	return &response, nil
}

// DeleteUserProfile ensures that the user profile is deleted correctly and does not exist already.
func (s *UserProfileService) DeleteUserProfile(userID uint) error {
	// Business validation: check if profile exists before attempting delete.
	_, err := s.repo.GetByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &ErrUserNotFound{UserID: userID}
		}
		return err
	}

	if err := s.repo.Delete(userID); err != nil {
		return err
	}
	return nil
}

package repository

import (
	"errors"
	"gorm.io/gorm"
	"profile-service/internal/models"
)

// IUserProfileRepository defines the methods required for user profile operations.
type IUserProfileRepository interface {
	Create(profile *models.UserProfile) error
	GetByUserID(userID uint) (*models.UserProfile, error)
	Update(profileID uint, profileData models.UserProfile) error
	Delete(userID uint) error
}

// UserProfileRepository handles database operations for user profiles using GORM.
type UserProfileRepository struct {
	db *gorm.DB
}

// NewUserProfileRepository returns a new instance of a GORM-based user profile repository.
func NewUserProfileRepository(db *gorm.DB) IUserProfileRepository {
	return &UserProfileRepository{db}
}

// Create adds a new UserProfile to the database and handles possible errors.
func (repo *UserProfileRepository) Create(profile *models.UserProfile) error {
	if profile == nil {
		return errors.New("profile is nil")
	}
	if err := repo.db.Create(profile).Error; err != nil {
		return err
	}
	return nil
}

// GetByUserID finds a UserProfile by its associated user ID and handles not found errors.
func (repo *UserProfileRepository) GetByUserID(userID uint) (*models.UserProfile, error) {
	var profile models.UserProfile
	result := repo.db.Where("user_id = ?", userID).First(&profile)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &profile, nil
}

// Update modifies an existing UserProfile by profileID with new profile data and handles potential errors.
func (repo *UserProfileRepository) Update(profileID uint, profileData models.UserProfile) error {
	result := repo.db.Model(&models.UserProfile{}).Where("id = ?", profileID).Updates(profileData)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected, the profile may not exist")
	}
	return nil
}

// Delete removes a UserProfile from the database by its associated user ID.
func (repo *UserProfileRepository) Delete(userID uint) error {
	result := repo.db.Where("user_id = ?", userID).Delete(&models.UserProfile{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows affected, the profile may not exist")
	}
	return nil
}

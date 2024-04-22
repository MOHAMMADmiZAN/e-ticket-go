package repository

import (
	"auth-service/internal/models"
	"errors"
	"gorm.io/gorm"
)

type IUserVerificationRepository interface {
	CreateVerification(verification *models.UserVerification) error
	UpdateVerification(verification *models.UserVerification) error
	FindVerificationByID(id uint) (*models.UserVerification, error)
	FindVerificationsByUserID(userID uint) ([]*models.UserVerification, error)
}
type UserVerificationRepository struct {
	DB *gorm.DB
}

func NewUserVerificationRepository(db *gorm.DB) IUserVerificationRepository {
	return &UserVerificationRepository{DB: db}
}

// CreateVerification adds a new user verification entry to the database.
func (r *UserVerificationRepository) CreateVerification(verification *models.UserVerification) error {
	if result := r.DB.Create(verification); result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateVerification modifies an existing user verification entry in the database.
func (r *UserVerificationRepository) UpdateVerification(verification *models.UserVerification) error {
	if result := r.DB.Save(verification); result.Error != nil {
		return result.Error
	}
	return nil
}

// FindVerificationByID retrieves a user verification by its ID.
func (r *UserVerificationRepository) FindVerificationByID(id uint) (*models.UserVerification, error) {
	var verification models.UserVerification
	result := r.DB.First(&verification, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user verification not found")
		}
		return nil, result.Error
	}
	return &verification, nil
}

// FindVerificationsByUserID retrieves all verifications for a specific user from the database.
func (r *UserVerificationRepository) FindVerificationsByUserID(userID uint) ([]*models.UserVerification, error) {
	var verifications []*models.UserVerification
	result := r.DB.Where("user_id = ?", userID).Find(&verifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return verifications, nil
}

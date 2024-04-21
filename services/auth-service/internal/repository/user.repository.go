package repository

import (
	"auth-service/internal/models"
	"gorm.io/gorm"
)

// IUserRepository provides an interface for database operations involving users.
type IUserRepository interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	ExistsByUsernameOrEmail(username, email string) (bool, error)
	UpdatePassword(userID uint, newPassword string) error
	Delete(userID uint) error
}

// UserRepository is a GORM-based implementation of IUserRepository.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

// Create inserts a new User into the database.
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByUsername finds a user by their username.
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail finds a user by their email.
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ExistsByUsernameOrEmail checks if a user exists with the given username or email.
func (r *UserRepository) ExistsByUsernameOrEmail(username, email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("username = ? OR email = ?", username, email).Count(&count).Error
	return count > 0, err
}

// UpdatePassword updates a user's password.
func (r *UserRepository) UpdatePassword(userID uint, newPassword string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("password", newPassword).Error
}

// Delete removes a user from the database.
func (r *UserRepository) Delete(userID uint) error {
	return r.db.Delete(&models.User{}, userID).Error
}

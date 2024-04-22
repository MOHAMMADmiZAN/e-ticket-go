package repository

import (
	"auth-service/internal/models"
	"time"

	"gorm.io/gorm"
)

// ILoginHistoryRepository defines the interface for login history repository operations.
type ILoginHistoryRepository interface {
	RecordLoginAttempt(history models.LoginHistory) error
	GetLoginAttempts(userID uint, from, to time.Time) ([]models.LoginHistory, error)
	GetHistoryByUserID(userID uint) ([]models.LoginHistory, error)
	UpdateLoginHistory(history *models.LoginHistory) error
	UpdateLogoutTime(historyID uint, logoutTime time.Time) error
}

// LoginHistoryRepository is a concrete implementation of ILoginHistoryRepository.
type LoginHistoryRepository struct {
	db *gorm.DB
}

// NewLoginHistoryRepository returns a new instance of a LoginHistoryRepository.
func NewLoginHistoryRepository(db *gorm.DB) ILoginHistoryRepository {
	return &LoginHistoryRepository{db: db}
}

// RecordLoginAttempt saves a new login attempt to the database.
func (repo *LoginHistoryRepository) RecordLoginAttempt(history models.LoginHistory) error {
	return repo.db.Create(&history).Error
}

// GetLoginAttempts retrieves all login attempts for a user within a specified time range.
func (repo *LoginHistoryRepository) GetLoginAttempts(userID uint, from, to time.Time) ([]models.LoginHistory, error) {
	var histories []models.LoginHistory
	err := repo.db.Where("user_id = ? AND login_time BETWEEN ? AND ?", userID, from, to).Find(&histories).Error
	return histories, err
}

// GetHistoryByUserID retrieves all login attempts for a specific user.
func (repo *LoginHistoryRepository) GetHistoryByUserID(userID uint) ([]models.LoginHistory, error) {
	var histories []models.LoginHistory
	err := repo.db.Where("user_id = ?", userID).Find(&histories).Error
	return histories, err
}

// UpdateLoginHistory updates an existing login history entry in the database.
func (repo *LoginHistoryRepository) UpdateLoginHistory(history *models.LoginHistory) error {
	return repo.db.Save(history).Error
}

// UpdateLogoutTime updates the logout time for a specific login history entry.
func (repo *LoginHistoryRepository) UpdateLogoutTime(historyID uint, logoutTime time.Time) error {
	return repo.db.Model(&models.LoginHistory{}).Where("id = ?", historyID).Update("logout_time", logoutTime).Error
}

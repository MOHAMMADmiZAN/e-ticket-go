package services

import (
	"auth-service/internal/api/dto"
	"auth-service/internal/repository"
	"time"
)

// Define thresholds for suspicious activity
const (
	SuspiciousActivityThreshold = 5                // Number of failed attempts to trigger warning
	SuspiciousActivityTimeFrame = 10 * time.Minute // Time frame for counting failed attempts
)

type ILoginHistoryService interface {
	RecordLoginAttempt(req dto.CreateLoginHistoryRequest) (dto.LoginHistoryResponse, error)
	GetLoginAttemptsByUser(userID uint, from, to time.Time) ([]dto.LoginHistoryResponse, error)
	UpdateLoginHistory(historyID uint, req dto.UpdateLoginHistoryRequest) error
	RecordUserLogout(historyID uint, logoutTime time.Time) error
	CheckSuspiciousActivity(userID uint) (bool, error)
}

type LoginHistoryService struct {
	loginHistoryRepo repository.ILoginHistoryRepository
}

func NewLoginHistoryService(loginHistoryRepo repository.ILoginHistoryRepository) ILoginHistoryService {
	return &LoginHistoryService{
		loginHistoryRepo: loginHistoryRepo,
	}
}

func (s *LoginHistoryService) RecordLoginAttempt(req dto.CreateLoginHistoryRequest) (dto.LoginHistoryResponse, error) {
	history := req.ToHistoryModel()
	err := s.loginHistoryRepo.RecordLoginAttempt(history)
	if err != nil {
		return dto.LoginHistoryResponse{}, err
	}
	response := dto.FromHistoryModel(history)
	return response, nil
}

func (s *LoginHistoryService) GetLoginAttemptsByUser(userID uint, from, to time.Time) ([]dto.LoginHistoryResponse, error) {
	histories, err := s.loginHistoryRepo.GetLoginAttempts(userID, from, to)
	if err != nil {
		return nil, err
	}
	var responses []dto.LoginHistoryResponse
	for _, history := range histories {
		responses = append(responses, dto.FromHistoryModel(history))
	}
	return responses, nil
}

func (s *LoginHistoryService) UpdateLoginHistory(historyID uint, req dto.UpdateLoginHistoryRequest) error {
	history := req.ToUpdateHistoryModel()
	history.ID = historyID
	return s.loginHistoryRepo.UpdateLoginHistory(&history)
}

func (s *LoginHistoryService) RecordUserLogout(historyID uint, logoutTime time.Time) error {
	// Update the logout time for the given history ID
	return s.loginHistoryRepo.UpdateLogoutTime(historyID, logoutTime)
}

func (s *LoginHistoryService) CheckSuspiciousActivity(userID uint) (bool, error) {
	// Fetch login attempts within the suspicious activity time frame
	from := time.Now().Add(-SuspiciousActivityTimeFrame)
	to := time.Now()
	histories, err := s.loginHistoryRepo.GetLoginAttempts(userID, from, to)
	if err != nil {
		return false, err
	}
	// Count the number of failed attempts
	var failedAttempts int
	for _, history := range histories {
		if !history.Successful {
			failedAttempts++
		}
	}
	// Check if failed attempts to exceed the threshold
	isSuspicious := failedAttempts >= SuspiciousActivityThreshold
	return isSuspicious, nil
}

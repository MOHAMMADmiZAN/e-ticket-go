package services

import (
	"auth-service/internal/api/dto"
	"auth-service/internal/repository"
	"time"
)

type IUserVerificationService interface {
	CreateVerification(req dto.CreateUserVerificationRequest) (dto.UserVerificationResponse, error)
	UpdateVerificationStatus(id uint, status string, verifiedAt *time.Time) (dto.UserVerificationResponse, error)
	GetVerificationDetails(id uint) (dto.UserVerificationResponse, error)
	GetVerificationsByUserID(userID uint) ([]dto.UserVerificationResponse, error)
}

type UserVerificationService struct {
	repo repository.IUserVerificationRepository
}

// NewUserVerificationService Constructor function to initialize a new UserVerificationService with its dependencies.
func NewUserVerificationService(repo repository.IUserVerificationRepository) IUserVerificationService {
	return &UserVerificationService{
		repo: repo,
	}
}

// CreateVerification handles the logic to create a new user verification entry.
func (s *UserVerificationService) CreateVerification(req dto.CreateUserVerificationRequest) (dto.UserVerificationResponse, error) {
	verification := req.ToVerificationModel()
	if err := s.repo.CreateVerification(&verification); err != nil {
		return dto.UserVerificationResponse{}, err
	}
	return dto.FromVerificationModel(verification), nil
}

// UpdateVerificationStatus updates the status and verified time of an existing verification record.
func (s *UserVerificationService) UpdateVerificationStatus(id uint, status string, verifiedAt *time.Time) (dto.UserVerificationResponse, error) {
	verification, err := s.repo.FindVerificationByID(id)
	if err != nil {
		return dto.UserVerificationResponse{}, err
	}

	verification.VerificationStatus = status
	if verifiedAt != nil && !verification.VerifiedAt.Equal(*verifiedAt) {
		verification.VerifiedAt = *verifiedAt
	}

	if err = s.repo.UpdateVerification(verification); err != nil {
		return dto.UserVerificationResponse{}, err
	}

	return dto.FromVerificationModel(*verification), nil
}

// GetVerificationDetails retrieves the details of a verification by its ID.
func (s *UserVerificationService) GetVerificationDetails(id uint) (dto.UserVerificationResponse, error) {
	verification, err := s.repo.FindVerificationByID(id)
	if err != nil {
		return dto.UserVerificationResponse{}, err
	}
	return dto.FromVerificationModel(*verification), nil
}

// GetVerificationsByUserID retrieves all verification records for a given user.
func (s *UserVerificationService) GetVerificationsByUserID(userID uint) ([]dto.UserVerificationResponse, error) {
	verifications, err := s.repo.FindVerificationsByUserID(userID)
	if err != nil {
		return nil, err
	}
	var responses []dto.UserVerificationResponse
	for _, verification := range verifications {
		responses = append(responses, dto.FromVerificationModel(*verification))
	}
	return responses, nil
}

package dto

import (
	"auth-service/internal/models"
	"auth-service/pkg"
	"time"
)

type CreateUserVerificationRequest struct {
	UserID           uint   `json:"userID" binding:"required"`
	VerificationType string `json:"verificationType" binding:"required"`
}

func (v *CreateUserVerificationRequest) ToVerificationModel() models.UserVerification {
	token, err := pkg.GenerateToken()
	if err != nil {
		panic(err)
	}
	return models.UserVerification{
		UserID:             v.UserID,
		VerificationType:   v.VerificationType,
		VerificationStatus: "pending",
		VerificationToken:  token,
		ExpirationDate:     time.Now().Add(24 * time.Hour), // Expiration typically set for 24 hours later
	}
}

type UserVerificationResponse struct {
	VerificationID     uint      `json:"verificationID"`
	UserID             uint      `json:"userID"`
	VerificationType   string    `json:"verificationType"`
	VerificationStatus string    `json:"verificationStatus"`
	VerificationToken  string    `json:"verificationToken"`
	ExpirationDate     time.Time `json:"expirationDate"`
	VerifiedAt         time.Time `json:"verifiedAt"`
}

func FromVerificationModel(v models.UserVerification) UserVerificationResponse {
	return UserVerificationResponse{
		VerificationID:     v.ID,
		UserID:             v.UserID,
		VerificationType:   v.VerificationType,
		VerificationStatus: v.VerificationStatus,
		VerificationToken:  v.VerificationToken,
		ExpirationDate:     v.ExpirationDate,
		VerifiedAt:         v.VerifiedAt,
	}
}

type UpdateUserVerificationRequest struct {
	VerificationStatus string    `json:"verificationStatus" binding:"required,oneof=pending verified failed"`
	VerifiedAt         time.Time `json:"verifiedAt" binding:"omitempty"`
}

func (u *UpdateUserVerificationRequest) ToUpdateVerificationModel() models.UserVerification {
	var verification models.UserVerification
	verification.VerificationStatus = u.VerificationStatus
	if !u.VerifiedAt.IsZero() {
		verification.VerifiedAt = u.VerifiedAt
	}
	return verification
}

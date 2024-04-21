package dto

import (
	"auth-service/internal/models"
	"time"
)

type CreateLoginHistoryRequest struct {
	UserID            uint   `json:"userID" binding:"required"`
	IPAddress         string `json:"ipAddress" binding:"required,ip"`
	DeviceInformation string `json:"deviceInformation" binding:"required"`
	Successful        bool   `json:"successful"`
	FailureReason     string `json:"failureReason,omitempty"`
}

func (l *CreateLoginHistoryRequest) ToHistoryModel() models.LoginHistory {
	return models.LoginHistory{
		UserID:            l.UserID,
		LoginTime:         time.Now(),  // Set login time during creation
		LogoutTime:        time.Time{}, // Optional, set on logout
		IPAddress:         l.IPAddress,
		DeviceInformation: l.DeviceInformation,
		Successful:        l.Successful,
		FailureReason:     l.FailureReason,
	}
}

type LoginHistoryResponse struct {
	LoginHistoryID    uint      `json:"loginHistoryID"`
	UserID            uint      `json:"userID"`
	LoginTime         time.Time `json:"loginTime"`
	LogoutTime        time.Time `json:"logoutTime"`
	IPAddress         string    `json:"ipAddress"`
	DeviceInformation string    `json:"deviceInformation"`
	Successful        bool      `json:"successful"`
	FailureReason     string    `json:"failureReason"`
}

func FromHistoryModel(l models.LoginHistory) LoginHistoryResponse {
	return LoginHistoryResponse{
		LoginHistoryID:    l.ID,
		UserID:            l.UserID,
		LoginTime:         l.LoginTime,
		LogoutTime:        l.LogoutTime,
		IPAddress:         l.IPAddress,
		DeviceInformation: l.DeviceInformation,
		Successful:        l.Successful,
		FailureReason:     l.FailureReason,
	}
}

type UpdateLoginHistoryRequest struct {
	LogoutTime    time.Time `json:"logoutTime" binding:"omitempty"`
	Successful    bool      `json:"successful" binding:"omitempty"`
	FailureReason string    `json:"failureReason" binding:"omitempty"`
}

func (l *UpdateLoginHistoryRequest) ToUpdateHistoryModel() models.LoginHistory {
	var loginHistory models.LoginHistory
	if !l.LogoutTime.IsZero() {
		loginHistory.LogoutTime = l.LogoutTime
	}
	loginHistory.Successful = l.Successful
	if l.FailureReason != "" {
		loginHistory.FailureReason = l.FailureReason
	}
	return loginHistory
}

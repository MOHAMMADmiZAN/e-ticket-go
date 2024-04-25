package dto

import (
	"notification-service/internal/models"
	"time"
)

// CreateNotificationRequest defines the structure for a request to create a notification.
type CreateNotificationRequest struct {
	UserID       uint      `json:"userID" validate:"required"`
	Type         string    `json:"type" validate:"required"`
	Status       string    `json:"status" validate:"required"`
	Channel      string    `json:"channel" validate:"required"`
	Content      string    `json:"content" validate:"required"`
	SendDate     time.Time `json:"sendDate" validate:"required"`
	Acknowledged bool      `json:"acknowledged"`
}

// ToModel converts CreateNotificationRequest to Notification model.
func (req *CreateNotificationRequest) ToModel() models.Notification {
	return models.Notification{
		UserID:       req.UserID,
		Type:         req.Type,
		Status:       req.Status,
		Channel:      req.Channel,
		Content:      req.Content,
		SendDate:     req.SendDate,
		Acknowledged: req.Acknowledged,
	}
}

// NotificationResponse defines the structure for a response that includes notification data.
type NotificationResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"userID"`
	Type         string    `json:"type"`
	Status       string    `json:"status"`
	Channel      string    `json:"channel"`
	Content      string    `json:"content"`
	SendDate     time.Time `json:"sendDate"`
	Acknowledged bool      `json:"acknowledged"`
}

// FromModel converts Notification model to NotificationResponse.
func FromModel(model models.Notification) *NotificationResponse {
	return &NotificationResponse{
		ID:           model.ID,
		UserID:       model.UserID,
		Type:         model.Type,
		Status:       model.Status,
		Channel:      model.Channel,
		Content:      model.Content,
		SendDate:     model.SendDate,
		Acknowledged: model.Acknowledged,
	}
}

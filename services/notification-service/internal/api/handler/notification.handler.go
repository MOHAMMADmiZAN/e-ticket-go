package handler

import (
	"fmt"
	"net/http"
	"notification-service/internal/api/dto"
	"notification-service/internal/services"
	"notification-service/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NotificationHandler handles HTTP requests related to notifications.
type NotificationHandler struct {
	notificationService services.INotificationService
}

// NewNotificationHandler creates a new instance of NotificationHandler.
func NewNotificationHandler(notificationService services.INotificationService) *NotificationHandler {
	return &NotificationHandler{
		notificationService: notificationService,
	}
}

// CreateNotification handles POST /notifications
// @Summary Create notification
// @Description Creates a new notification based on the provided details.
// @Tags notifications
// @Accept json
// @Produce json
// @Param notification body dto.CreateNotificationRequest true "Create Notification Request"
// @Success 201 {object} dto.NotificationResponse "Notification created successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid notification data"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router / [post]
func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var req dto.CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid notification data: %v", err))
		return
	}

	response, err := h.notificationService.CreateNotification(req)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusCreated, response, "Notification created successfully")
}

// DeleteNotification handles DELETE /notifications/{id}
// @Summary Delete notification
// @Description Deletes a specific notification by ID.
// @Tags notifications
// @Produce json
// @Param id path int true "Notification ID"
// @Success 204 "Notification deleted successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid notification ID"
// @Failure 404 {object} pkg.APIResponse "Notification not found"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /{id} [delete]
func (h *NotificationHandler) DeleteNotification(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid notification ID: %v", err))
		return
	}

	if err := h.notificationService.DeleteNotification(uint(id)); err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusNoContent, nil, "Notification deleted successfully")
}

// GetNotification handles GET /notifications/{id}
// @Summary Get notification
// @Description Retrieves a specific notification by ID.
// @Tags notifications
// @Produce json
// @Param id path int true "Notification ID"
// @Success 200 {object} dto.NotificationResponse "Notification fetched successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid notification ID"
// @Failure 404 {object} pkg.APIResponse "Notification not found"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /{id} [get]
func (h *NotificationHandler) GetNotification(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid notification ID: %v", err))
		return
	}

	response, err := h.notificationService.GetNotification(uint(id))
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, response, "Notification fetched successfully")
}

// ListNotifications handles GET /notifications
// @Summary List notifications
// @Description Retrieves a list of notifications, optionally filtered by user ID.
// @Tags notifications
// @Produce json
// @Param userID query int false "User ID"
// @Success 200 {array} dto.NotificationResponse "Notifications fetched successfully"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router / [get]
func (h *NotificationHandler) ListNotifications(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Query("userID"), 10, 32) // optional user ID filter

	offset, _ := strconv.Atoi(c.Query("offset")) // pagination offset
	limit, _ := strconv.Atoi(c.Query("limit"))   // pagination limit

	responses, err := h.notificationService.ListNotifications(uint(userID), offset, limit)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, responses, "Notifications fetched successfully")
}

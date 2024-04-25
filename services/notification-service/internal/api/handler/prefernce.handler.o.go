package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"notification-service/internal/api/dto"
	"notification-service/internal/services"
	"notification-service/pkg"
	"strconv"
)

type UserNotificationHandler struct {
	userNotificationService services.UserNotificationService
}

func NewUserNotificationHandler(service services.UserNotificationService) *UserNotificationHandler {
	return &UserNotificationHandler{
		userNotificationService: service,
	}
}

// CreateUserPreferences handles POST /users/{userID}/preferences
// @Summary Create user notification preferences
// @Description Creates notification preferences for a specific user.
// @Tags User Preferences
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Param body body dto.UserNotificationPreferencesRequest true "Create Notification Preferences Request"
// @Success 201 {object} dto.UserNotificationPreferencesResponse "Preferences created successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid input data"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /users/{userID}/preferences [post]
func (h *UserNotificationHandler) CreateUserPreferences(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	var req dto.UserNotificationPreferencesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid input data: %v", err))
		return
	}

	err = h.userNotificationService.CreateUserPreferences(uint(userID), req)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusCreated, req, "Preferences created successfully")
}

// UpdateUserPreferences handles PUT /users/{userID}/preferences
// @Summary Update user notification preferences
// @Description Updates notification preferences for a specific user.
// @Tags User Preferences
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Param body body dto.UserNotificationPreferencesUpdate true "Update Notification Preferences Request"
// @Success 200 {object} dto.UserNotificationPreferencesResponse "Preferences updated successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid input data"
// @Failure 404 {object} pkg.APIResponse "Preferences not found"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /users/{userID}/preferences [put]
func (h *UserNotificationHandler) UpdateUserPreferences(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	var update dto.UserNotificationPreferencesUpdate
	if err := c.ShouldBindJSON(&update); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid input data: %v", err))
		return
	}

	err = h.userNotificationService.UpdateUserPreferences(uint(userID), update)
	if err != nil {
		pkg.RespondWithError(c, determineStatusCode(err), err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, update, "Preferences updated successfully")
}

// GetUserPreferences handles GET /users/{userID}/preferences
// @Summary Get user notification preferences
// @Description Retrieves notification preferences for a specific user.
// @Tags User Preferences
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} dto.UserNotificationPreferencesResponse "Preferences fetched successfully"
// @Failure 404 {object} pkg.APIResponse "Preferences not found"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /users/{userID}/preferences [get]
func (h *UserNotificationHandler) GetUserPreferences(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	response, err := h.userNotificationService.GetUserPreferences(uint(userID))
	if err != nil {
		pkg.RespondWithError(c, determineStatusCode(err), err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, response, "Preferences fetched successfully")
}

// DeleteUserPreferences handles DELETE /users/{userID}/preferences
// @Summary Delete user notification preferences
// @Description Deletes notification preferences for a specific user.
// @Tags User Preferences
// @Produce json
// @Param userID path int true "User ID"
// @Success 204 "Preferences deleted successfully"
// @Failure 404 {object} pkg.APIResponse "Preferences not found"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /users/{userID}/preferences [delete]
func (h *UserNotificationHandler) DeleteUserPreferences(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	if err := h.userNotificationService.DeleteUserPreferences(uint(userID)); err != nil {
		pkg.RespondWithError(c, determineStatusCode(err), err)
		return
	}

	c.Status(http.StatusNoContent)
}

// Helper function to determine the HTTP status code based on the error type.
func determineStatusCode(err error) int {
	//if err == services.ErrNotFound {
	//	return http.StatusNotFound
	//} else if err == services.ErrInvalidInput {
	//	return http.StatusBadRequest
	//} else if err == services.ErrUnauthorized {
	//	return http.StatusUnauthorized
	//} else if err == services.ErrConflict {
	//	return http.StatusConflict
	//}
	return http.StatusInternalServerError // Default if unspecified
}

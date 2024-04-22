package handler

import (
	"auth-service/internal/api/dto"
	"auth-service/internal/services"
	"auth-service/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type LoginHistoryHandler struct {
	loginHistoryService services.ILoginHistoryService
}

func NewLoginHistoryHandler(loginHistoryService services.ILoginHistoryService) *LoginHistoryHandler {
	return &LoginHistoryHandler{
		loginHistoryService: loginHistoryService,
	}
}

// RecordLoginAttempt handles POST /login-attempts
// @Summary Record login attempt
// @Description This endpoint records a login attempt for a user.
// @Tags login-history
// @Accept json
// @Produce json
// @Param loginAttempt body dto.CreateLoginHistoryRequest true "Create Login History Request"
// @Success 201 {object} pkg.APIResponse "Login attempt recorded successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid login attempt data"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /login-attempts [post]
func (h *LoginHistoryHandler) RecordLoginAttempt(c *gin.Context) {
	var req dto.CreateLoginHistoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid login attempt data: %v", err))
		return
	}

	response, err := h.loginHistoryService.RecordLoginAttempt(req)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusCreated, response, "Login attempt recorded successfully")
}

// GetLoginAttemptsByUser handles GET /login-attempts/{userID}
// @Summary Get login attempts by user
// @Description This endpoint fetches all login attempts for a specified user within a given time frame.
// @Tags login-history
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Param from query string true "Start time (RFC3339 format)"
// @Param to query string true "End time (RFC3339 format)"
// @Success 200 {object} pkg.APIResponse "Login attempts fetched successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid parameters"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /login-attempts/{userID} [get]
func (h *LoginHistoryHandler) GetLoginAttemptsByUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	from, err := time.Parse(time.RFC3339, c.Query("from"))
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid from time: %v", err))
		return
	}

	to, err := time.Parse(time.RFC3339, c.Query("to"))
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid to time: %v", err))
		return
	}

	responses, err := h.loginHistoryService.GetLoginAttemptsByUser(uint(userID), from, to)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, responses, "Login attempts fetched successfully")
}

// RecordUserLogout handles PUT /logout/{historyID}
// @Summary Record user logout
// @Description This endpoint records the logout time for a specific login history record.
// @Tags login-history
// @Accept json
// @Produce json
// @Param historyID path int true "History ID"
// @Success 200 {object} pkg.APIResponse "Logout recorded successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid history ID"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /logout/{historyID} [post]
func (h *LoginHistoryHandler) RecordUserLogout(c *gin.Context) {
	historyID, err := strconv.ParseUint(c.Param("historyID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid history ID: %v", err))
		return
	}

	logoutTime := time.Now() // Assuming logout at current time
	if err := h.loginHistoryService.RecordUserLogout(uint(historyID), logoutTime); err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, nil, "Logout recorded successfully")
}

// CheckSuspiciousActivity handles GET /suspicious-activity/{userID}
// @Summary Check for suspicious activity
// @Description This endpoint checks for suspicious login activity for a specified user.
// @Tags login-history
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} pkg.APIResponse "Suspicious activity status returned"
// @Failure 400 {object} pkg.APIResponse "Invalid user ID"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /suspicious-activity/{userID} [get]
func (h *LoginHistoryHandler) CheckSuspiciousActivity(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	isSuspicious, err := h.loginHistoryService.CheckSuspiciousActivity(uint(userID))
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, gin.H{"is_suspicious": isSuspicious}, "Suspicious activity check completed")
}

package handler

import (
	"fmt"
	"net/http"
	"profile-service/internal/api/dto"
	"profile-service/internal/services"
	"profile-service/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService services.IUserProfileService
}

func NewProfileHandler(profileService services.IUserProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

// CreateProfile handles POST /create
// @Summary Create user profile
// @Description Creates a new user profile with the provided details.
// @Tags profile
// @Accept json
// @Produce json
// @Param profile body dto.UserProfileRequest true "Create Profile Request"
// @Success 201 {object} dto.UserProfileResponse "Profile created successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid profile data"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /create [post]
func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	var req dto.UserProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid profile data: %v", err))
		return
	}

	profile, err := h.profileService.CreateUserProfile(req)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusCreated, profile, "Profile created successfully")
}

// GetProfile handles GET /{userID}
// @Summary Get user profile
// @Description Retrieves the user profile for a specified user ID.
// @Tags profile
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} dto.UserProfileResponse "Profile fetched successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid user ID"
// @Failure 404 {object} pkg.APIResponse "Profile not found"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /{userID} [get]
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	profile, err := h.profileService.GetUserProfile(uint(userID))
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, profile, "Profile fetched successfully")
}

// UpdateProfile handles PUT /{userID}
// @Summary Update user profile
// @Description Updates the profile details for a given user ID.
// @Tags profile
// @Accept json
// @Produce json
// @Param userID path int true "User ID"
// @Param profile body dto.UserProfileUpdate true "Update Profile Request"
// @Success 200 {object} dto.UserProfileResponse "Profile updated successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid data"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /{userID} [put]
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	var req dto.UserProfileUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid update data: %v", err))
		return
	}

	profile, err := h.profileService.UpdateUserProfile(uint(userID), req)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, profile, "Profile updated successfully")
}

// DeleteProfile handles DELETE /{userID}
// @Summary Delete user profile
// @Description Deletes the profile of the specified user ID.
// @Tags profile
// @Produce json
// @Param userID path int true "User ID"
// @Success 204 "Profile deleted successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid user ID"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /{userID} [delete]
func (h *ProfileHandler) DeleteProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID: %v", err))
		return
	}

	if err := h.profileService.DeleteUserProfile(uint(userID)); err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusNoContent, nil, "Profile deleted successfully")
}

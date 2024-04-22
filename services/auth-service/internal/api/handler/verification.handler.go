package handler

import (
	"auth-service/internal/api/dto"
	"auth-service/internal/services"
	"auth-service/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserVerificationHandler struct {
	verificationService services.IUserVerificationService
}

func NewUserVerificationHandler(verificationService services.IUserVerificationService) *UserVerificationHandler {
	return &UserVerificationHandler{
		verificationService: verificationService,
	}
}

// CreateVerification handles POST /verifications
// @Summary Create user verification
// @Description This endpoint creates a user verification entry.
// @Tags verification
// @Accept json
// @Produce json
// @Param verification body dto.CreateUserVerificationRequest true "Create User Verification Request"
// @Success 201 {object} pkg.APIResponse "Verification created successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid request data"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /verifications [post]
func (h *UserVerificationHandler) CreateVerification(c *gin.Context) {
	var req dto.CreateUserVerificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid request data: %v", err))
		return
	}
	response, err := h.verificationService.CreateVerification(req)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}
	pkg.RespondWithSuccess(c, http.StatusCreated, response, "Verification created successfully")
}

// UpdateVerificationStatus handles PUT /verifications/{id}
// @Summary Update verification status
// @Description This endpoint updates the status of an existing verification.
// @Tags verification
// @Accept json
// @Produce json
// @Param id path int true "Verification ID"
// @Param status body dto.UpdateUserVerificationRequest true "Update Verification Status Request"
// @Success 200 {object} pkg.APIResponse "Verification status updated successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid verification ID or request data"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /verifications/{id} [put]
func (h *UserVerificationHandler) UpdateVerificationStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid verification ID: %v", err))
		return
	}
	var req dto.UpdateUserVerificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid request data: %v", err))
		return
	}
	response, err := h.verificationService.UpdateVerificationStatus(uint(id), req.VerificationStatus, &req.VerifiedAt)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}
	pkg.RespondWithSuccess(c, http.StatusOK, response, "Verification status updated successfully")
}

// GetVerificationDetails handles GET /verifications/{id}
// @Summary Get verification details
// @Description This endpoint retrieves the details of a specific verification.
// @Tags verification
// @Produce json
// @Param id path int true "Verification ID"
// @Success 200 {object} pkg.APIResponse "Verification details retrieved successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid verification ID"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /verifications/{id} [get]
func (h *UserVerificationHandler) GetVerificationDetails(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid verification ID: %v", err))
		return
	}
	response, err := h.verificationService.GetVerificationDetails(uint(id))
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}
	pkg.RespondWithSuccess(c, http.StatusOK, response, "Verification details retrieved successfully")
}

// GetVerificationsByUserID handles GET /users/{userID}/verifications
// @Summary Get all verifications for a user
// @Description This endpoint retrieves all verification entries associated with a specified user.
// @Tags verification
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} pkg.APIResponse "List of verifications fetched successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid user ID"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /users/{userID}/verifications [get]
func (h *UserVerificationHandler) GetVerificationsByUserID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID: %v", err))
		return
	}
	responses, err := h.verificationService.GetVerificationsByUserID(uint(userID))
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, fmt.Errorf("failed to fetch verifications: %v", err))
		return
	}
	pkg.RespondWithSuccess(c, http.StatusOK, responses, "List of verifications fetched successfully")
}

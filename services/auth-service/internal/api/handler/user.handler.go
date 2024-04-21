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

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// RegisterUser handles POST /auth/register endpoint
// @Summary Register new user
// @Description This endpoint registers a new user with the provided credentials.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "Create User Request"
// @Success 201 {object} pkg.APIResponse "User created successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid user data"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /register [post]
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var userDTO dto.CreateUserRequest
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user data: %v", err))
		return
	}

	userResponse, err := h.userService.RegisterUser(userDTO)
	if err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusCreated, userResponse, "User created successfully")
}

// AuthenticateUser handles POST /auth/login endpoint
// @Summary Authenticate user
// @Description This endpoint authenticates a user using username and password.
// @Tags auth
// @Accept json
// @Produce json
// @Param login body dto.UserLoginDTO true "User Login DTO"
// @Success 200 {object} pkg.APIResponse "User authenticated successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid login data"
// @Failure 401 {object} pkg.APIResponse "Unauthorized - Invalid credentials"
// @Router /login [post]
func (h *UserHandler) AuthenticateUser(c *gin.Context) {
	var loginDTO dto.UserLoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid login data: %v", err))
		return
	}

	response, err := h.userService.AuthenticateUser(loginDTO)
	if err != nil {
		pkg.RespondWithError(c, http.StatusUnauthorized, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, response, "User authenticated successfully")
}

// UpdateUserPassword handles PUT /users/{id}/password endpoint
// @Summary Update user password
// @Description This endpoint updates the password for the user with the specified ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param password body dto.UserPasswordUpdateDTO true "User Password Update DTO"
// @Success 200 {object} pkg.APIResponse "Password updated successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid user ID format or password data"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /users/{id}/password [put]
func (h *UserHandler) UpdateUserPassword(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID format: %v", err))
		return
	}

	var passwordDTO dto.UserPasswordUpdateDTO
	if err := c.ShouldBindJSON(&passwordDTO); err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid password data: %v", err))
		return
	}
	passwordDTO.UserID = uint(userID)

	if err := h.userService.UpdateUserPassword(passwordDTO); err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusOK, nil, "Password updated successfully")
}

// DeleteUser handles DELETE /users/{id} endpoint
// @Summary Delete user
// @Description This endpoint deletes a user with the specified ID.
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 204 {object} pkg.APIResponse "User deleted successfully"
// @Failure 400 {object} pkg.APIResponse "Invalid user ID format"
// @Failure 500 {object} pkg.APIResponse "Internal server error"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		pkg.RespondWithError(c, http.StatusBadRequest, fmt.Errorf("invalid user ID format: %v", err))
		return
	}

	if err := h.userService.DeleteUser(uint(userID)); err != nil {
		pkg.RespondWithError(c, http.StatusInternalServerError, err)
		return
	}

	pkg.RespondWithSuccess(c, http.StatusNoContent, nil, "User deleted successfully")
}

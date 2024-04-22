package dto

import (
	"auth-service/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=3,max=255"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=admin customer"`
}

func (c *CreateUserRequest) ToUserModel() models.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	return models.User{
		Username:            c.Username,
		Email:               c.Email,
		Password:            string(hashedPassword),
		Role:                c.Role,
		AccountCreationDate: time.Now(),
		Verified:            false,
	}
}

type UserResponse struct {
	UserID              uint      `json:"userID"`
	Username            string    `json:"username"`
	Email               string    `json:"email"`
	Role                string    `json:"role"`
	AccountCreationDate time.Time `json:"accountCreationDate"`
	Verified            bool      `json:"verified"`
}

func FromUserModel(u models.User) UserResponse {
	return UserResponse{
		UserID:              u.ID,
		Username:            u.Username,
		Email:               u.Email,
		Role:                u.Role,
		AccountCreationDate: u.AccountCreationDate,
		Verified:            u.Verified,
	}
}

type UpdateUserRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=6"`
	Role     string `json:"role" binding:"omitempty,oneof=admin customer"`
	Verified bool   `json:"verified" binding:"omitempty"`
}

func (u *UpdateUserRequest) ToUpdateUserModel() models.User {
	var user models.User
	if u.Email != "" {
		user.Email = u.Email
	}
	if u.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
	}
	if u.Role != "" {
		user.Role = u.Role
	}
	return user
}

type UserLoginDTO struct {
	Username string `json:"username" binding:"required,alphanum,min=3,max=255"`
	Password string `json:"password" binding:"required,min=6"`
}

// UserPasswordUpdateDTO represents the data required to update a user's password.
type UserPasswordUpdateDTO struct {
	UserID   uint   `json:"userID" binding:"required"` // Ensure UserID is passed securely, not from client directly in a real-world scenario
	Password string `json:"password" binding:"required,min=6"`
}

// ToUpdatePasswordModel Function to update a user's password in the user model
func (dto *UserPasswordUpdateDTO) ToUpdatePasswordModel() models.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	return models.User{
		Password: string(hashedPassword),
	}
}

// UserLoginResponseDto  includes only the user's ID and JWT token for authenticated sessions.
type UserLoginResponseDto struct {
	UserID uint   `json:"userID"`
	Token  string `json:"token"`
}

// UserLoginResponse  creates a new instance of UserLoginResponse.
func UserLoginResponse(userID uint, token string) *UserLoginResponseDto {
	return &UserLoginResponseDto{
		UserID: userID,
		Token:  token,
	}
}

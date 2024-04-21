package services

import (
	"auth-service/internal/api/dto"
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type IUserService interface {
	RegisterUser(userDTO dto.CreateUserRequest) (*dto.UserResponse, error)
	AuthenticateUser(loginDTO dto.UserLoginDTO) (*dto.UserLoginResponseDto, error)
	UpdateUserPassword(passwordDTO dto.UserPasswordUpdateDTO) error
	DeleteUser(userID uint) error
}

var JwtSecretKey = os.Getenv("JWT_SECRET")

type UserService struct {
	userRepo repository.IUserRepository
}

type Claims struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// RegisterUser handles the registration logic for a new user.
func (s *UserService) RegisterUser(userDTO dto.CreateUserRequest) (*dto.UserResponse, error) {
	exists, err := s.userRepo.ExistsByUsernameOrEmail(userDTO.Username, userDTO.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("a user with the given username or email already exists")
	}

	user := userDTO.ToUserModel()
	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}

	response := dto.FromUserModel(user)
	return &response, nil
}

// AuthenticateUser verifies the username and password and returns the user.
func (s *UserService) AuthenticateUser(loginDTO dto.UserLoginDTO) (*dto.UserLoginResponseDto, error) {
	user, err := s.userRepo.FindByUsername(loginDTO.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Compare the provided password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := s.GenerateJWT(*user)
	if err != nil {
		return nil, err
	}

	return dto.UserLoginResponse(user.ID, token), nil
}

// UpdateUserPassword updates a user's password.
func (s *UserService) UpdateUserPassword(passwordDTO dto.UserPasswordUpdateDTO) error {
	// Ensure the new password is hashed in the DTO
	user := passwordDTO.ToUpdatePasswordModel()
	user.ID = passwordDTO.UserID // Attach the UserID to the model
	return s.userRepo.UpdatePassword(user.ID, user.Password)
}

// DeleteUser removes a user from the database.
func (s *UserService) DeleteUser(userID uint) error {
	return s.userRepo.Delete(userID)
}

// GenerateJWT generates a new JWT token for a given user.
func (s *UserService) GenerateJWT(user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

package models

import (
	"gorm.io/gorm"
	"time"
)

// User defines a user model with related fields and relationships.
type User struct {
	gorm.Model                             // Embedding gorm.Model already gives you an auto-incrementing ID, created_at, updated_at, deleted_at.
	Username            string             `gorm:"uniqueIndex;size:255" json:"username"` // Unique username
	Email               string             `gorm:"uniqueIndex;size:255" json:"email"`    // Unique email
	Password            string             `json:"password"`
	Role                string             `json:"role"` // User role
	AccountCreationDate time.Time          `json:"accountCreationDate"`
	UserVerifications   []UserVerification `json:"userVerifications"` // One-to-Many relationship
	LoginHistories      []LoginHistory     `json:"loginHistories"`    // One-to-Many relationship
}

// TableName overrides the table name used by User to `users`.
func (User) TableName() string {
	return "users"
}

// LoginHistory defines a login history model with related fields and a belongs-to relationship with User.
type LoginHistory struct {
	gorm.Model                  // Embedding gorm.Model gives you an auto-incrementing ID, created_at, updated_at, deleted_at.
	UserID            uint      `json:"userId"` // Foreign key for User
	LoginTime         time.Time `json:"loginTime"`
	LogoutTime        time.Time `json:"logoutTime"`
	IPAddress         string    `json:"ipAddress"`
	DeviceInformation string    `json:"deviceInformation"`
	Successful        bool      `json:"successful"`
	FailureReason     string    `json:"failureReason"`
	User              User      `gorm:"foreignKey:UserID"` // Belongs to User
}

// TableName overrides the table name used by LoginHistory to `login_histories`.
func (LoginHistory) TableName() string {
	return "login_histories"
}

// UserVerification defines a user verification model with related fields and a belongs-to relationship with User.
type UserVerification struct {
	gorm.Model                   // Embedding gorm.Model gives you an auto-incrementing ID, created_at, updated_at, deleted_at.
	UserID             uint      `json:"userId"` // Foreign key for User
	VerificationType   string    `json:"verificationType"`
	VerificationStatus string    `json:"verificationStatus"`
	VerificationToken  string    `json:"verificationToken"`
	ExpirationDate     time.Time `json:"expirationDate"`
	VerifiedAt         time.Time `json:"verifiedAt"`
	User               User      `gorm:"foreignKey:UserID"` // Belongs to User
}

// TableName overrides the table name used by UserVerification to `user_verifications`.
func (UserVerification) TableName() string {
	return "user_verifications"
}

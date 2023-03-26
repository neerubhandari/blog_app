package models

import (
	"gorm.io/gorm"
)

// embed the gorm.Model struct, which gives us fields for ID, CreatedAt, UpdatedAt, and DeletedAt
type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

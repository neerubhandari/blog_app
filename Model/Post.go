package models

import "gorm.io/gorm"

// embed the gorm.Model struct, which gives us fields for ID, CreatedAt, UpdatedAt, and DeletedAt
type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

package controllers

import (
	models "blog_app/Model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PostController is a controller for managing blog posts
type PostController struct {
	DB *gorm.DB
}

// GetAllPosts returns a list of all blog posts
func (p *PostController) GetAllPosts(c *gin.Context) {
	var posts []models.Post
	p.DB.Find(&posts)

	c.JSON(200, gin.H{"data": posts})
}

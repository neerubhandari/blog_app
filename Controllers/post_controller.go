package controllers

import (
	models "blog_app/Model"
	"net/http"

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
	result := p.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve posts"})
		return
	}

	if len(posts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No posts found"})
		return
	}
	c.JSON(200, gin.H{"data": posts})
}

// create post
func CreatePost(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": post})
}

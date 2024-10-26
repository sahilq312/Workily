package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/model"
)

// AddLike adds a like to a post by a specific user
func AddLike(c *gin.Context) {
	var body struct {
		UserID uint `json:"user_id"`
		PostID uint `json:"post_id"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if like already exists
	var existingLike model.Like
	if initializer.DB.Where("user_id = ? AND post_id = ?", body.UserID, body.PostID).First(&existingLike).Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already liked"})
		return
	}

	like := model.Like{UserID: body.UserID, PostID: body.PostID}
	if err := initializer.DB.Create(&like).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add like"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Like added successfully"})
}

// RemoveLike removes a like from a post by a specific user
func RemoveLike(c *gin.Context) {
	var body struct {
		UserID uint `json:"user_id"`
		PostID uint `json:"post_id"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := initializer.DB.Where("user_id = ? AND post_id = ?", body.UserID, body.PostID).Delete(&model.Like{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like removed successfully"})
}

// GetLikesOnPost retrieves all likes on a specific post
func GetLikesOnPost(c *gin.Context) {
	postID := c.Param("post_id")

	var likes []model.Like
	if err := initializer.DB.Where("post_id = ?", postID).Find(&likes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch likes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"likes": likes})
}

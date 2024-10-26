package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/model"
)

// AddComment adds a comment to a post
func AddComment(c *gin.Context) {
	var body struct {
		Content string `json:"content"`
		UserID  uint   `json:"user_id"`
		PostID  uint   `json:"post_id"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	comment := model.Comment{
		Content: body.Content,
		UserID:  body.UserID,
		PostID:  body.PostID,
	}
	if err := initializer.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully", "comment": comment})
}

// GetCommentsOnPost retrieves all comments on a specific post
func GetCommentsOnPost(c *gin.Context) {
	postID := c.Param("post_id")

	var comments []model.Comment
	if err := initializer.DB.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

// DeleteComment deletes a comment by its ID
func DeleteComment(c *gin.Context) {
	commentID := c.Param("comment_id")

	if err := initializer.DB.Delete(&model.Comment{}, commentID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

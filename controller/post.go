package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/model"
)

func CreatePost(c *gin.Context) {
	//Get request body
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		UserID  uint   `json:"user_id"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}
	//Create post
	post := model.Post{
		Title:   body.Title,
		Content: body.Content,
		UserID:  body.UserID,
	}
	result := initializer.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Can store the post",
		})
		return
	}
	//Return post
	c.JSON(http.StatusCreated, gin.H{
		"job": result,
	})

}

func GetPosts(c *gin.Context) {
	id := c.Param("id")
	post := model.Post{}
	result := initializer.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	post := model.Post{}
	result := initializer.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	post := model.Post{}
	result := initializer.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}
	var body struct { // Define the structure of the request body
		Title   string `json:"title"` // Title of the job
		Content string `json:"content"`
	}
	err := c.BindJSON(&body) // Bind the request body to the defined structure
	if err != nil {          // If there's an error while binding the request body
		c.JSON(http.StatusBadRequest, gin.H{ // Return a bad request response
			"error": "Invalid request body", // Error message
		})
		return
	}

	// Update post
	initializer.DB.Model(&model.Post{}).Where("id = ?", id).Updates(body) // Update the post with the specified ID
	// Return post
	c.JSON(http.StatusOK, gin.H{ // If the post is updated successfully, return a success response
		"message": "Post updated successfully", // Success message
	})
}

func DeletePost(c *gin.Context) {
	//Get post
	id := c.Param("id")
	post := model.Post{}
	result := initializer.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}
	//Delete post
	result = initializer.DB.Delete(&post)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Can delete the post",
		})
		return
	}
	//Return post
	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}

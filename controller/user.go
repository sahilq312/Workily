package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	c.Get("user")
	var body struct {
		Name     string
		Email    string
		Password string
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request",
		})
		return
	}
	// Get request body

	//result := initializer.DB.Update()
}

func DeleteUser(c *gin.Context) {
	// Get request body
	c.Get("user")

}

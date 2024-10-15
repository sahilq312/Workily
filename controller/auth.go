package controller

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	//Get the request body

	//Check if user Exist

	//Compare Hashed Password

	//Set session

	//Return User and session

}

func Register(c *gin.Context) {
	//Get the request body

	//Check if user already exist

	//hash the password

	//set Session

	//Return User and session
}

func GetUser(c *gin.Context) {
	//Get User from Session
	
	c.JSON(200, gin.H{
		"user": "user",
	})
	//Return User
}

func Logout(c *gin.Context) {
	//Delete Session

	//Return Success
}

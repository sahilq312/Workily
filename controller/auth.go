package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/model"
	"github.com/sahilq312/workly/utils"
)

// Login function to authenticate a user
func Login(c *gin.Context) {
	// type of input
	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// valid request
	var body loginRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// Find user
	var user model.User
	result := initializer.DB.Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not exist"})
		return
	}
	// match password
	match, err := utils.CompareHashedPassword(body.Password, user.Password)
	if err != nil || !match {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	// Set jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret not set"})
		return
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in generating JWT token"})
		return
	}
	// set cookie to browser
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", false, true)

	// Sending User as Response
	c.JSON(http.StatusOK, gin.H{
		"data": user, // Consider returning a sanitized user object
	})
}

// Register function to register a new user
func Register(c *gin.Context) {
	//struct for request
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//Get the request body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Request"})
		return
	}

	//Check if user already exist
	var userExist model.User
	if initializer.DB.Where("email = ?", body.Email).First(&userExist).Error == nil && userExist.ID != 0 {
		c.JSON(400, gin.H{
			"error": "User already exists",
		})
		return
	}

	//hash the password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error in hashing the password",
		})
		return
	}

	//Create a new user
	user := model.User{Name: body.Name, Email: body.Email, Password: hashedPassword}
	result := initializer.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "Error in creating the user",
		})
	}
	// Set jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret not set"})
		return
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in generating JWT token"})
		return
	}
	// set cookie to browser
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", false, true)
	//Return User and session
	c.JSON(200, gin.H{
		"data": user,
	})
}

func GetUserById(c *gin.Context) {
	// Get id from params
	id := c.Params.ByName("id")

	//find user with the given id
	var user model.User
	if err := initializer.DB.Where("id = ?", id).First(&user).Error; err != nil {
		// If no user is found, return a 404 error
		c.JSON(404, gin.H{
			"error": "No user found",
		})
		return
	}
	// return user
	c.JSON(200, gin.H{
		"data": user,
	})
}

func GetUser(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "No user found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func Logout(c *gin.Context) {
	//Delete Session
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	//Return Success
	c.JSON(http.StatusOK, gin.H{
		"data": "Logged out successfully",
	})
}

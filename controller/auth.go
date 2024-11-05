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
	// Define the structure for the login request
	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind the request body to the loginRequest structure
	var body loginRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Find the user by email
	var user model.User
	result := initializer.DB.Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not exist"})
		return
	}

	// Compare the provided password with the user's password
	match, err := utils.CompareHashedPassword(body.Password, user.Password)
	if err != nil || !match {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	})

	// Retrieve the JWT secret from environment variables
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT secret not set"})
		return
	}

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in generating JWT token"})
		return
	}

	// Set the JWT token as a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", false, true)

	// Return the user details as a response
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

// Register function to register a new user
func Register(c *gin.Context) {
	// Define the structure for the registration request
	type registerRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind the request body to the registerRequest structure
	var body registerRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	// Check if the user already exists
	var userExist model.User
	if initializer.DB.Where("email = ?", body.Email).First(&userExist).Error == nil && userExist.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error in hashing the password",
		})
		return
	}

	// Create a new user
	user := model.User{Name: body.Name, Email: body.Email, Password: hashedPassword}
	result := initializer.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error in creating the user",
		})
		return
	}

	// Generate JWT token for the new user
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

	// Set the JWT token as a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", false, true)

	// Return the user details and session
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func GetUserById(c *gin.Context) {
	// Extract the user ID from the URL parameters
	id := c.Params.ByName("id")

	// Find the user by ID
	var user model.User
	if err := initializer.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No user found",
		})
		return
	}

	// Return the user details
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func GetUser(c *gin.Context) {
	// Retrieve the user from the context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "No user found",
		})
		return
	}
	// Return the user details
	userData, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cast user to model.User",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":    userData.ID,
			"name":  userData.Name,
			"email": userData.Email,
		},
	})
}

func Logout(c *gin.Context) {
	// Delete the JWT token cookie
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	// Return success message
	c.JSON(http.StatusOK, gin.H{
		"data": "Logged out successfully",
	})
}

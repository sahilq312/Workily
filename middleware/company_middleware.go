package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/model"
)

func CompanyAuth(c *gin.Context) {
	tokenString, err := c.Cookie("CompanyAuth")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "CompanyAuth cookie not found"})
		c.Abort()
		return
	}

	// Parse and validate JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_COMPANY_SECRET")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}

	// Check expiration
	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
		c.Abort()
		return
	}

	// Retrieve and validate company
	var company model.Company
	initializer.DB.First(&company, claims["company_id"])
	if company.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Company not found"})
		c.Abort()
		return
	}
	c.Set("company", company)
	c.Next()
}

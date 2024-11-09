package controller

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/model"
	"github.com/sahilq312/workly/utils"
	"gorm.io/gorm"
)

// CreateCompany creates a new company
func CreateCompany(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Logo     string `json:"logo"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Address  string `json:"address"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if email for company exists
	var companyExists model.Company
	if err := initializer.DB.Where("email = ?", body.Email).First(&companyExists).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exists",
		})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	// Create the company record
	company := model.Company{
		Name:     body.Name,
		Logo:     body.Logo,
		Email:    body.Email,
		Password: hashedPassword,
		Address:  body.Address,
	}
	if err := initializer.DB.Create(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in creating the company"})
		return
	}

	// Return the created company without the password
	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":      company.ID,
			"name":    company.Name,
			"logo":    company.Logo,
			"email":   company.Email,
			"address": company.Address,
		},
	})
}

func LoginCompany(c *gin.Context) {
	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body loginRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	// Check if company exists
	var company model.Company
	result := initializer.DB.Where("email = ?", body.Email).First(&company)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Company does not exist"})
		return
	}

	// Validate password
	match, err := utils.CompareHashedPassword(body.Password, company.Password)
	if err != nil || !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"company_id": company.ID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
		"iat":        time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_COMPANY_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT token"})
		return
	}

	// Set the cookie with token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("CompanyAuth", tokenString, 86400, "/", "", false, true) // Cookie duration matches token expiration

	c.JSON(http.StatusOK, gin.H{
		"data": company,
	})
}

func GetCompany(c *gin.Context) {
	company, ok := c.Get("company")
	if !ok || company == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Company not found"})
		return
	}
	var companyModel = company.(model.Company)
	if companyModel.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": companyModel,
	})
}

// GetCompany retrieves a company by ID along with its jobs
func GetCompanyById(c *gin.Context) {
	// Parse company ID from URL and ensure it's valid
	companyID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	var company model.Company
	// Retrieve company by ID
	if err := initializer.DB.First(&company, uint(companyID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get company"})
		return
	}

	// Return the company data
	c.JSON(http.StatusOK, gin.H{"company": company})
}

// UpdateCompany updates a company's information
func UpdateCompany(c *gin.Context) {
	company, ok := c.Get("company")
	if !ok || company == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Company not found"})
		return
	}
	companyModel := company.(model.Company)
	if companyModel.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	var body struct {
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		Email   string `json:"email"`
		Address string `json:"address"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Update company fields
	if body.Name != "" {
		companyModel.Name = body.Name
	}
	if body.Logo != "" {
		companyModel.Logo = body.Logo
	}
	if body.Email != "" {
		companyModel.Email = body.Email
	}
	if body.Address != "" {
		companyModel.Address = body.Address
	}

	if result := initializer.DB.Save(&companyModel); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully", "company": companyModel})
}

// DeleteCompany deletes a company by ID
func DeleteCompany(c *gin.Context) {
	company, ok := c.Get("company")
	if !ok || company == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Company not found"})
		return
	}
	companyModel := company.(model.Company)
	if companyModel.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	if err := initializer.DB.Delete(&companyModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Company deleted successfully"})
}

// GetAllCompanies retrieves all companies along with their jobs
func GetAllCompanies(c *gin.Context) {
	var companies []model.Company
	// Fetch companies from the database
	if err := initializer.DB.Find(&companies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get companies"})
		return
	}
	// Respond with the list of companies (empty list if none found)
	c.JSON(http.StatusOK, gin.H{"data": companies})
}

func GetCompanyJobs(c *gin.Context) {
	company, ok := c.Get("company")
	if !ok || company == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Company not found"})
		return
	}
	companyModel := company.(model.Company)
	if companyModel.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	var jobs []model.Job
	if err := initializer.DB.Where("company_id = ?", companyModel.ID).Find(&jobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get jobs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": jobs})

}

func GetCompanyJobById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}
	var job model.Job
	if err := initializer.DB.First(&job, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": job})
}

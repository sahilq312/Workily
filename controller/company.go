package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/model"
	"gorm.io/gorm"
)

// CreateCompany creates a new company
func CreateCompany(c *gin.Context) {
	var body struct {
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		OwnerID uint   `json:"owner_id"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate request body
	if body.Name == "" || body.Logo == "" || body.OwnerID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Check if owner exists
	var owner model.User
	if err := initializer.DB.First(&owner, body.OwnerID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Owner not found"})
		return
	}

	// Create company
	company := model.Company{Name: body.Name, Logo: body.Logo, OwnerID: body.OwnerID}
	if result := initializer.DB.Create(&company); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Company created successfully", "company": company})
}

// GetCompany retrieves a company by ID along with its jobs
func GetCompany(c *gin.Context) {
	companyID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	var company model.Company
	result := initializer.DB.Preload("Jobs").First(&company, companyID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"company": company})
}

// UpdateCompany updates a company's information
func UpdateCompany(c *gin.Context) {
	companyID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	var body struct {
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		OwnerID uint   `json:"owner_id"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if owner exists
	var owner model.User
	if body.OwnerID != 0 {
		if err := initializer.DB.First(&owner, body.OwnerID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Owner not found"})
			return
		}
	}

	var company model.Company
	result := initializer.DB.First(&company, companyID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	company.Name = body.Name
	company.Logo = body.Logo
	company.OwnerID = body.OwnerID

	if result := initializer.DB.Save(&company); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully", "company": company})
}

// DeleteCompany deletes a company by ID
func DeleteCompany(c *gin.Context) {
	companyID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	var company model.Company
	result := initializer.DB.First(&company, companyID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	if err := initializer.DB.Delete(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}

// GetAllCompanies retrieves all companies along with their jobs
func GetAllCompanies(c *gin.Context) {
	var companies []model.Company
	result := initializer.DB.Preload("Jobs").Find(&companies)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get companies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"companies": companies})
}

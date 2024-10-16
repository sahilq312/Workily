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
	// Get request body
	var body struct {
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		OwnerID uint   `json:"owner_id"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// Validate request body
	if body.Name == "" || body.Logo == "" || body.OwnerID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}

	// Create company
	company := model.Company{Name: body.Name, Logo: body.Logo, OwnerID: body.OwnerID}
	result := initializer.DB.Create(&company)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create company",
		})
		return
	}

	// Return company
	c.JSON(http.StatusCreated, gin.H{
		"message": "Company created successfully",
		"company": company,
	})
}

// GetCompany retrieves a company by ID
func GetCompany(c *gin.Context) {
	companyID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid company ID",
		})
		return
	}

	var company model.Company
	result := initializer.DB.First(&company, companyID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Company not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get company",
		})
		return
	}

	// Return company
	c.JSON(http.StatusOK, company)
}

// UpdateCompany updates a company
func UpdateCompany(c *gin.Context) {
	companyID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid company ID",
		})
		return
	}

	var body struct {
		Name    string `json:"name"`
		Logo    string `json:"logo"`
		OwnerID uint   `json:"owner_id"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if body.Name == "" || body.Logo == "" || body.OwnerID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}

	var company model.Company
	result := initializer.DB.First(&company, companyID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Company not found",
		})
		return
	}

	company.Name = body.Name
	company.Logo = body.Logo
	company.OwnerID = body.OwnerID

	result = initializer.DB.Save(&company)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update company",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Company updated successfully",
		"company": company,
	})
}

// DeleteCompany deletes a company
func DeleteCompany(c *gin.Context) {
	companyID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid company ID",
		})
		return
	}

	var company model.Company
	result := initializer.DB.Delete(&company, companyID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete company",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Company deleted successfully",
	})
}

// GetAllCompanies retrieves all companies
func GetAllCompanies(c *gin.Context) {
	var companies []model.Company
	result := initializer.DB.Find(&companies)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get companies",
		})
		return
	}

	c.JSON(http.StatusOK, companies)
}

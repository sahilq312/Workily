package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/model"
)

func ApplyForJob(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	userModel, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user type assertion"})
		return
	}
	userID := userModel.ID
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	var body struct {
		JobID uint `json:"job_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if body.JobID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Job ID is required"})
		return
	}

	application := model.Application{
		UserID: userID,
		JobID:  body.JobID,
	}
	if result := initializer.DB.Create(&application); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit application"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Application submitted successfully"})
}

func GetUserApplications(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	userModel, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user type assertion"})
		return
	}
	userID := userModel.ID
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	var applications []model.Application
	if result := initializer.DB.Where("user_id = ?", userID).Find(&applications); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch applications"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"applications": applications})
}

func DeleteApplicationByCompany(c *gin.Context) {
	company, ok := c.Get("company")
	if !ok || company == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Company not found"})
		return
	}
	companyModel, ok := company.(model.Company)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid company type assertion"})
		return
	}
	companyID := companyModel.ID
	if companyID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid company ID"})
		return
	}
	jobID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job ID"})
		return
	}
	result := initializer.DB.Where("company_id = ?", companyID).Delete(&model.Application{}, jobID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete application"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}

func DeleteApplication(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	userModel, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user type assertion"})
		return
	}
	userID := userModel.ID
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}
	applicationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}
	result := initializer.DB.Where("user_id = ?", userID).Delete(&model.Application{}, applicationID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete application"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}

func GetApplicationByID(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	userModel, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user type assertion"})
		return
	}
	userID := userModel.ID
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}
	applicationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}
	var application model.Application
	if result := initializer.DB.Where("user_id = ?", userID).First(&application, applicationID); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch application"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"application": application})
}

func GetApplicationsByCompany(c *gin.Context) {
	company, ok := c.Get("company")
	if !ok || company == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Company not found"})
		return
	}
	companyModel, ok := company.(model.Company)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid company type assertion"})
		return
	}
	companyID := companyModel.ID
	if companyID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid company ID"})
		return
	}
	var applications []model.Application
	if result := initializer.DB.Where("company_id = ?", companyID).Find(&applications); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch applications"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"applications": applications})
}

func UpdateApplicationStatusByCompany(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}
	var body struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if body.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
		return
	}
	company, ok := c.Get("company")
	if !ok || company == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Company not found"})
		return
	}
	companyModel, ok := company.(model.Company)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid company type assertion"})
		return
	}
	companyID := companyModel.ID
	if companyID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid company ID"})
		return
	}
	// Check if the application exists and belongs to the company before updating
	var application model.Application
	if result := initializer.DB.Where("company_id = ?", companyID).Where("id = ?", id).First(&application); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found or does not belong to this company"})
		return
	}
	result := initializer.DB.Model(&model.Application{}).Where("company_id = ?", companyID).Where("id = ?", id).Update("status", body.Status)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update application status"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Application status updated successfully"})
}

package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/model"
)

func CreateJob(c *gin.Context) {
	var body struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Location    string   `json:"location"`
		Salary      string   `json:"salary"`
		CompanyID   uint     `json:"company_id"`
		Skills      []string `json:"skills"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if body.Title == "" || body.Description == "" || body.Location == "" || body.Salary == "" || body.CompanyID == 0 || len(body.Skills) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "All fields are required",
		})
		return
	}

	job := model.Job{
		Title:       body.Title,
		Description: body.Description,
		Location:    body.Location,
		Salary:      body.Salary,
		CompanyID:   body.CompanyID,
		Skills:      body.Skills,
	}
	result := initializer.DB.Create(&job)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create job",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"job": job,
	})
}

func GetJob(c *gin.Context) {
	id := c.Param("id")
	job := model.Job{}
	result := initializer.DB.First(&job, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Job not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"job": job,
	})
}

func UpdateJob(c *gin.Context) {
	var body struct {
		Title       string
		Description string
		Skills      []string
		Location    string
		Salary      string
		CompanyId   string
		jobId       string
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	job := model.Job{}
	result := initializer.DB.First(&job, body.jobId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Job not found",
		})
		return
	}

	job.Title = body.Title
	job.Description = body.Description
	job.Skills = body.Skills
	job.Location = body.Location
	job.Salary = body.Salary

	result = initializer.DB.Save(&job)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update job",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Job updated successfully",
	})
}

func DeleteJob(c *gin.Context) {
	id := c.Param("id")
	job := model.Job{}
	result := initializer.DB.First(&job, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Job not found",
		})
		return
	}

	result = initializer.DB.Delete(&job)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete job",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Job deleted successfully",
	})
}

func GetAllJobs(c *gin.Context) {
	jobs := []model.Job{}
	initializer.DB.Find(&jobs)
	c.JSON(http.StatusOK, gin.H{
		"jobs": jobs,
	})
}

func GetJobsByCompany(c *gin.Context) {
	companyId := c.Param("company_id")
	jobs := []model.Job{}
	initializer.DB.Where("company_id = ?", companyId).Find(&jobs)
	c.JSON(http.StatusOK, gin.H{
		"jobs": jobs,
	})
}

func GetJobsByLocation(c *gin.Context) {
	var body struct {
		Location string `json:"location"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	jobs := []model.Job{}
	initializer.DB.Where("location = ?", body.Location).Find(&jobs)

	if len(jobs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No jobs found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"jobs": jobs,
	})
}

func GetJobsBySkill(c *gin.Context) {
	var body struct {
		Skill string `json:"skill"`
	}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	jobs := []model.Job{}
	initializer.DB.Where("skills LIKE ?", "%"+body.Skill+"%").Find(&jobs)

	if len(jobs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No jobs found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"jobs": jobs,
	})
}

package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/initializer"
	"github.com/sahilq312/workly/model"
)

// CreateJob creates a new job
func CreateJob(c *gin.Context) {
	// Retrieve company from context
	company, ok := c.Get("company")
	if !ok || company == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Company not found"})
		return
	}
	companyModel, ok := company.(model.Company)
	if !ok || companyModel.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	// Bind request body
	var body struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Location    string   `json:"location"`
		Salary      string   `json:"salary"`
		Skills      []string `json:"skills"` // Skill names
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate required fields
	if body.Title == "" || body.Description == "" || body.Location == "" || body.Salary == "" || len(body.Skills) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Fetch or create skills
	var skills []model.Skill
	for _, skillName := range body.Skills {
		var skill model.Skill
		if err := initializer.DB.FirstOrCreate(&skill, model.Skill{Name: skillName}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing skills"})
			return
		}
		skills = append(skills, skill)
	}

	// Create the job
	job := model.Job{
		Title:       body.Title,
		Description: body.Description,
		Location:    body.Location,
		Salary:      body.Salary,
		CompanyID:   companyModel.ID,
		Skills:      skills,
	}
	if err := initializer.DB.Create(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job"})
		return
	}

	// Return the created job
	c.JSON(http.StatusCreated, gin.H{"data": job})
}

// GetJob retrieves a job by ID
func GetJob(c *gin.Context) {
	id := c.Param("id")
	var job model.Job
	result := initializer.DB.Preload("Skills").First(&job, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"job": job})
}

// UpdateJob updates an existing job
func UpdateJob(c *gin.Context) {
	var body struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Location    string   `json:"location"`
		Salary      string   `json:"salary"`
		CompanyID   uint     `json:"company_id"`
		Skills      []string `json:"skills"` // Skill names
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	jobID := c.Param("id")
	var job model.Job
	result := initializer.DB.First(&job, jobID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	// Update job fields
	job.Title = body.Title
	job.Description = body.Description
	job.Location = body.Location
	job.Salary = body.Salary

	// Update skills
	var skills []model.Skill
	for _, skillName := range body.Skills {
		var skill model.Skill
		initializer.DB.FirstOrCreate(&skill, model.Skill{Name: skillName})
		skills = append(skills, skill)
	}
	job.Skills = skills

	if err := initializer.DB.Save(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job updated successfully"})
}

// DeleteJob deletes a job by ID
func DeleteJob(c *gin.Context) {
	company, _ := c.Get("company")
	id := c.Param("id")

	valid := initializer.DB.Model(&model.Job{}).Where("id = ? AND company_id = ?", id, company).First(&model.Job{}).Error
	if valid != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to delete this job"})
		return
	}

	var job model.Job
	result := initializer.DB.First(&job, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	if err := initializer.DB.Delete(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
}

// GetAllJobs retrieves all jobs
func GetAllJobs(c *gin.Context) {
	page := 1
	perPage := 3
	pageStr := c.Query("page")

	// Get filtering and search parameters
	title := c.Query("title")
	location := c.Query("location")
	skills := c.QueryArray("skills") // skills provided as an array of skill IDs
	search := c.Query("search")

	// Parse page number if provided
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	// Initialize query on Job model
	query := initializer.DB.Model(&model.Job{})

	// Apply filters if provided
	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}
	if location != "" {
		query = query.Where("location ILIKE ?", "%"+location+"%")
	}
	if len(skills) > 0 {
		// Join with skills table to filter jobs by skill IDs
		query = query.Joins("JOIN job_skills ON job_skills.job_id = jobs.id").
			Where("job_skills.skill_id IN ?", skills).
			Group("jobs.id")
	}
	if search != "" {
		// Apply search across title and description fields
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Count total rows after applying filters
	var totalRows int64
	query.Count(&totalRows)
	totalPages := (totalRows + int64(perPage) - 1) / int64(perPage)

	// Retrieve the jobs with pagination
	offset := (page - 1) * perPage
	var jobs []model.Job
	result := query.Offset(offset).Limit(perPage).Find(&jobs)
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "No response",
		})
		return
	}

	// Return paginated and filtered results
	c.JSON(http.StatusOK, gin.H{
		"jobs":       jobs,
		"totalPages": totalPages,
		"page":       page,
		"totalRows":  totalRows,
	})
}

// GetJobsByCompany retrieves jobs by company ID
func GetJobsByCompany(c *gin.Context) {
	companyId := c.Param("company_id")
	var jobs []model.Job
	initializer.DB.Preload("Skills").Where("company_id = ?", companyId).Find(&jobs)
	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}

// GetJobsByLocation retrieves jobs by location
func GetJobsByLocation(c *gin.Context) {
	var body struct {
		Location string `json:"location"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var jobs []model.Job
	initializer.DB.Preload("Skills").Where("location = ?", body.Location).Find(&jobs)
	if len(jobs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No jobs found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}

// GetJobsBySkill retrieves jobs by skill
func GetJobsBySkill(c *gin.Context) {
	var body struct {
		Skill string `json:"skill"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var jobs []model.Job
	// Join with skills table and filter jobs by the specified skill name
	err := initializer.DB.
		Joins("JOIN job_skills ON job_skills.job_id = jobs.id").
		Joins("JOIN skills ON skills.id = job_skills.skill_id").
		Where("skills.name = ?", body.Skill).
		Preload("Skills").
		Find(&jobs).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve jobs"})
		return
	}

	if len(jobs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No jobs found for the specified skill"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}

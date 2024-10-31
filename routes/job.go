package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
	"github.com/sahilq312/workly/middleware"
)

func JobRoutes(r *gin.Engine) {
	job := r.Group("/job")
	job.GET("/", controller.GetAllJobs)
	job.POST("/create", middleware.CompanyAuth, controller.CreateJob)
	job.GET("/get/:id", controller.GetJob)
	job.PUT("/update/:id", middleware.CompanyAuth, controller.UpdateJob)
	job.DELETE("/delete/:id", middleware.CompanyAuth, controller.DeleteJob)
}

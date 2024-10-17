package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
)

func JobRoutes(r *gin.Engine) {
	job := r.Group("/job")
	job.POST("/create", controller.CreateJob)
	job.GET("/get/:id", controller.GetJob)
	job.PUT("/update/:id", controller.UpdateJob)
	job.DELETE("/delete/:id", controller.DeleteJob)
}

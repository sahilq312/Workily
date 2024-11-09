package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
	"github.com/sahilq312/workly/middleware"
)
func ApplicationRoutes(r *gin.Engine) {
	// APPLICATION ROUTES
	application := r.Group("/application")
	// ROUTE FOR USERS TO APPLY FOR JOBS
	application.POST("/apply", middleware.RequireAuth, controller.ApplyForJob)
	// ROUTE FOR USERS TO GET THEIR APPLICATIONS
	application.GET("/user", middleware.RequireAuth, controller.GetUserApplications)
	// ROUTE FOR USERS TO GET A PARTICULAR APPLICATION
	application.GET("/:id", middleware.RequireAuth, controller.GetApplicationByID)
	// ROUTE FOR USERS TO DELETE THEIR APPLICATIONS
	application.DELETE("/:id", middleware.RequireAuth, controller.DeleteApplication)
	// ROUTE FOR COMPANIES TO GET APPLICATIONS
	application.GET("/company/:id", middleware.CompanyAuth, controller.GetApplicationsByCompany)
	// ROUTE FOR COMPANIES TO UPDATE THE STATUS OF APPLICATIONS
	application.PATCH("/company/:id/status", middleware.CompanyAuth, controller.UpdateApplicationStatusByCompany)
	// ROUTE FOR COMPANIES TO DELETE APPLICATIONS
	application.DELETE("/company/:id", middleware.CompanyAuth, controller.DeleteApplicationByCompany)
}

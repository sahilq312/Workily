package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
	"github.com/sahilq312/workly/middleware"
)

func CompanyRoutes(r *gin.Engine) {
	company := r.Group("/company")
	company.POST("/create", controller.CreateCompany)
	company.POST("/login", controller.LoginCompany)
	company.GET("/", middleware.CompanyAuth, controller.GetCompany)
	company.GET("/get/:id", controller.GetCompanyById)
	company.PUT("/update/:id", middleware.CompanyAuth, controller.UpdateCompany)
	company.DELETE("/delete/:id", middleware.CompanyAuth, controller.DeleteCompany)
	company.GET("/get-all-companies", controller.GetAllCompanies)
}

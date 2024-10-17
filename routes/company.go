package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
)

func CompanyRoutes(r *gin.Engine) {
	company := r.Group("/company")
	company.POST("/create", controller.CreateCompany)
	company.GET("/get/:id", controller.GetCompany)
	company.PUT("/update/:id", controller.UpdateCompany)
	company.DELETE("/delete/:id", controller.DeleteCompany)
	company.GET("/get-all-companies", controller.GetAllCompanies)
}

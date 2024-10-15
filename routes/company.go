package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
)

func CompanyRoutes(r *gin.Engine) {
	r.POST("/create-company", controller.CreateCompany)
	r.GET("/get-company/:id", controller.GetCompany)
	r.PUT("/update-company/:id", controller.UpdateCompany)
	r.DELETE("/delete-company/:id", controller.DeleteCompany)
	r.GET("/get-all-companies", controller.GetAllCompanies)
}

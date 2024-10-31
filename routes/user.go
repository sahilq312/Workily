package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
	"github.com/sahilq312/workly/middleware"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/user")
	user.GET("/get/:id", controller.GetUser)
	user.PUT("/update/:id", middleware.RequireAuth, controller.UpdateUser)
	user.DELETE("/delete/:id", middleware.RequireAuth, controller.DeleteUser)
}

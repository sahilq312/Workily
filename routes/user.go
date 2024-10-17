package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/user")
	user.GET("/get/:id", controller.GetUser)
	user.PUT("/update/:id", controller.UpdateUser)
	user.DELETE("/delete/:id", controller.DeleteUser)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/user")
	user.DELETE("/delete/:id", controller.DeleteUser)
	user.PUT("/update/:id", controller.UpdateUser)
}

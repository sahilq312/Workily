package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controller.Login)
	r.POST("/signup", controller.Register)
	r.GET("/get-user", controller.GetUser)
	r.GET("/logout", controller.Logout)
	r.GET("/getuser/:id", controller.GetUserById)
}

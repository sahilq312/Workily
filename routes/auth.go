package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
	"github.com/sahilq312/workly/middleware"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/login", controller.Login)
	auth.POST("/signup", controller.Register)
	auth.GET("/get-user", middleware.RequireAuth, controller.GetUser)
	auth.GET("/logout", controller.Logout)
	auth.GET("/getuser/:id", controller.GetUserById)
}

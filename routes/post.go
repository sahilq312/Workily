package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
	"github.com/sahilq312/workly/middleware"
)

func PostRoutes(r *gin.Engine) {
	post := r.Group("/post")
	post.POST("/create", middleware.RequireAuth, controller.CreatePost)
	post.GET("/", controller.GetPosts)
	post.GET("/get/:id", controller.GetPost)
	post.PUT("/update/:id", middleware.RequireAuth, controller.UpdatePost)
	post.DELETE("/delete/:id", middleware.RequireAuth, controller.DeletePost)
}

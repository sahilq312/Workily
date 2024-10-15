package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
)

func PostRoutes(r *gin.Engine) {
	r.POST("/create-post", controller.CreatePost)
	r.GET("/get-posts", controller.GetPosts)
	r.GET("/get-post/:id", controller.GetPost)
	r.PUT("/update-post/:id", controller.UpdatePost)
	r.DELETE("/delete-post/:id", controller.DeletePost)
}

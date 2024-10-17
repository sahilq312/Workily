package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
)

func PostRoutes(r *gin.Engine) {
	post := r.Group("/post")
	post.POST("/create", controller.CreatePost)
	post.GET("/get/:id", controller.GetPost)
	post.PUT("/update/:id", controller.UpdatePost)
	post.DELETE("/delete/:id", controller.DeletePost)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
)

func CommentRoutes(r *gin.Engine) {
	r.POST("/comment", controller.AddComment)
	r.DELETE("/comment/:comment_id", controller.DeleteComment)
	r.GET("/comments/:post_id", controller.GetCommentsOnPost)
}

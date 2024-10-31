package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
	"github.com/sahilq312/workly/middleware"
)

func CommentRoutes(r *gin.Engine) {
	r.POST("/comment", middleware.RequireAuth, controller.AddComment)
	r.DELETE("/comment/:comment_id", middleware.RequireAuth, controller.DeleteComment)
	r.GET("/comments/:post_id", middleware.RequireAuth, controller.GetCommentsOnPost)
}

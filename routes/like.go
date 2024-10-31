package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
	"github.com/sahilq312/workly/middleware"
)

func LikeRoutes(r *gin.Engine) {
	r.POST("/like", middleware.RequireAuth, controller.AddLike)
	r.DELETE("/like", middleware.RequireAuth, controller.RemoveLike)
	r.GET("/likes/:post_id", middleware.RequireAuth, controller.GetLikesOnPost)
}

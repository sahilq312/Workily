package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilq312/workly/controller"
)

func LikeRoutes(r *gin.Engine) {
	r.POST("/like", controller.AddLike)
	r.DELETE("/like", controller.RemoveLike)
	r.GET("/likes/:post_id", controller.GetLikesOnPost)
}

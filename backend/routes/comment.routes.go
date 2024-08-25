package routes

import (
	"github.com/ritankarsaha/backend/controllers"
	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine) {
	router.POST("/comments", controllers.CreateComment)
	router.GET("/posts/:post_id/comments", controllers.GetCommentsByPostID)
	router.DELETE("/comments/:id", controllers.DeleteComment)
}

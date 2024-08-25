package routes

import (
	"github.com/ritankarsaha/backend/controllers"
	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine) {
	router.POST("/comments", controllers.CreateComment)
	router.GET("/posts/:newsID/comments", controllers.GetCommentsByNewsID)
	router.DELETE("/comments/:commentID", controllers.DeleteComment)
}

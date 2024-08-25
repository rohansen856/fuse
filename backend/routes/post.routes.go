package routes

import (
	"github.com/ritankarsaha/backend/controllers"
	"github.com/gin-gonic/gin"
)

func PostRoutes(router *gin.Engine) {
	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPostByID)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
}

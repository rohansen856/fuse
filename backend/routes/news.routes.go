package routes

import (
	"github.com/ritankarsaha/backend/controllers"
	"github.com/gin-gonic/gin"
)

func NewsRoutes(router *gin.Engine) {
	router.POST("/news", controllers.CreateNews)
	router.GET("/news", controllers.GetNews)
	router.GET("/news/:id", controllers.GetNewsByID)
	router.PATCH("/news/:id", controllers.UpdateNews)
	router.DELETE("/news/:id", controllers.DeleteNews)
	router.GET("/news/category/:category",controllers.GetNewsByCategory)
}


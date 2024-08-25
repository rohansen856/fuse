package routes

import (
	"github.com/ritankarsaha/backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/user/:user_id", controllers.GetUser())
}

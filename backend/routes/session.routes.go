package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/ritankarsaha/backend/controllers"
)

func SessionRoutes(router *gin.Engine) {
	router.POST("/sessions", controllers.CreateSession())
	router.GET("/sessions/:session_id", controllers.GetSession())
	router.DELETE("/sessions/:session_id", controllers.DeleteSession())
}

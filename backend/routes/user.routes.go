package routes

import (
	"github.com/ritankarsaha/backend/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/user/:user_id", controllers.GetUser())            
	router.PATCH("/user/:user_id", controllers.UpdateUser())        
	router.POST("/user/register", controllers.RegisterUser())       
	router.DELETE("/user/:user_id", controllers.DeleteUser())       
	router.GET("/users", controllers.ListUsers())                  
	router.PATCH("/user/:user_id/role", controllers.ChangeUserRole()) 
	router.GET("/user/:user_id/profile", controllers.GetUserProfile())
}

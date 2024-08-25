package main

import (
	"github.com/ritankarsaha/backend/config"
	"github.com/ritankarsaha/backend/database"
	"github.com/ritankarsaha/backend/middleware"
	"github.com/ritankarsaha/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.InitDatabase()
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	routes.PostRoutes(router)
	routes.CommentRoutes(router)
	router.Run(":" + config.AppConfig.Port)
}

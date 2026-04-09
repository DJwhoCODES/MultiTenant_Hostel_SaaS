package main

import (
	"log"

	"github.com/djwhocodes/hostel_saas/config"
	"github.com/djwhocodes/hostel_saas/internal/database"
	"github.com/djwhocodes/hostel_saas/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	database.ConnectMongo()

	database.InitIndexes()

	r := gin.Default()

	routes.RegisterRoutes(r)

	log.Println("🚀 Server running on port", config.AppConfig.Port)
	r.Run(":" + config.AppConfig.Port)
}

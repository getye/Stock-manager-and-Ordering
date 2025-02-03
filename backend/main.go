package main

import (
	"backend/config"
	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/uploads", "./uploads")

	routes.SetupUserRoutes(router, db)
	routes.SetupProductRoutes(router, db)

	router.Run(":5000")
}

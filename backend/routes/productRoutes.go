package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/api/products", controllers.CreateProduct(db))
	router.GET("/api/products", controllers.GetProducts(db))
}

package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/api/products/add", controllers.AddProductHandler(db))
	router.GET("/api/products/view", controllers.GetProducts(db))
	router.PUT("/api/products/update/:id", controllers.UpdateProduct(db))
	router.DELETE("/api/products/delete/:id", controllers.DeleteProduct(db))

}

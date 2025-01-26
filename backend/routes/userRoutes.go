package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/api/users/add", controllers.CreateUser(db))
	router.GET("/api/users/view", controllers.GetUsers(db))
	router.PUT("/api/users/update/:id", controllers.UpdateUser(db))
	router.PUT("/api/users/update/status/:id", controllers.UpdateUserStatus(db))
	router.DELETE("/api/users/delete/:id", controllers.DeleteUser(db))
}

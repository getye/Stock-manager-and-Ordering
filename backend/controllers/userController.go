package controllers

import (
	"backend/models"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		er := c.ShouldBindJSON(&user)
		if er != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
			return
		}

		if user.ID == uuid.Nil {
			user.ID = uuid.New()
		}

		err := models.InsertUser(db, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully!",
			"user":    user,
		})
	}
}

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := models.GetUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var updatedUser models.User

		err := c.ShouldBindJSON(&updatedUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		user, err := models.UpdateUser(db, id, updatedUser)
		if err != nil {
			if err.Error() == "user not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with ID %s not found", id)})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			}
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := models.DeleteUser(db, id)
		if err != nil {
			if err.Error() == "user not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with ID %s not found", id)})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User successfully deleted"})
	}
}

func UpdateUserStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var userStatus models.User

		err := c.ShouldBindJSON(&userStatus)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		user, err := models.UpdateUserStatus(db, id, userStatus)
		if err != nil {
			if err.Error() == "user not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with ID %s not found", id)})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			}
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

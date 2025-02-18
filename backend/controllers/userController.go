package controllers

import (
	"backend/models"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func generateOtp(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}

	hashedOtp, err := bcrypt.GenerateFromPassword([]byte(result), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error:", err)
		return ""
	}
	return string(hashedOtp)
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		er := c.ShouldBindJSON(&user)
		if er != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
			return
		}

		user.ID = uuid.New()
		user.Password = generateOtp(8)

		err := models.InsertUser(db, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// send OTP over user email here

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully!",
			"user":    user,
		})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest LoginRequest
		err := c.ShouldBindJSON(&loginRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		var user models.User
		result := db.Where("email = ?", loginRequest.Email).First(&user)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if user.Password != loginRequest.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		var redirect string

		if user.Role == "Admin" {
			redirect = "/user"
		} else if user.Role == "Suplier" {
			redirect = "/product"
		} else {
			redirect = "/login"
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Login successful",
			"user":     user,
			"redirect": redirect,
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

package controllers

import (
	"backend/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AddProductHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use Form to read other form fields
		name := c.DefaultPostForm("name", "")
		description := c.DefaultPostForm("description", "")
		quantityStr := c.DefaultPostForm("quantity", "")
		priceStr := c.DefaultPostForm("price", "")

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
			return
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price"})
			return
		}

		// Handle file upload
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
			return
		}

		timeStr := time.Now()
		layout := "15:04:05.9999999"

		t, err := time.Parse(layout, timeStr.Format(layout))
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}

		file_name := fmt.Sprintf("%02d%02d%02d%09d", t.Hour(), t.Minute(), t.Second(), t.Nanosecond())

		dst := fmt.Sprintf("uploads/%s.png", file_name)

		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}

		product := models.Product{
			ID:          uuid.New(),
			User_ID:     uuid.New(), // receive user id from params
			Name:        name,
			Description: description,
			Quantity:    quantity,
			Price:       price,
			Image:       dst,
		}

		err = models.AddProduct(db, &product)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Product added successfully!",
			"product": product,
		})
	}
}

func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := models.GetProducts(db)
		println("products: ", products)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"products": products})
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var updatedProduct models.Product

		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			log.Println("error", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		product, err := models.UpdateProduct(db, id, updatedProduct)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := models.DeleteProduct(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Product successfully deleted"})
	}
}

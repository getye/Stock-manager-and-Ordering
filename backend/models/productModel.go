package models

import (
	"fmt"
	"log"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	User_ID     uuid.UUID `gorm:"foreignKey;type:uuid;" json:"user_id"`
	Name        string    `gorm:"size:255;not null" json:"name"`
	Description string    `gorm:"size:255;not null" json:"description"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	Price       float64   `gorm:"not null" json:"price"`
	Order       int       `gorm:"not null" json:"order"`
	Image       string    `gorm:"size:255;not null" json:"image"`
}

func AddProduct(db *gorm.DB, product *Product) error {

	result := db.Create(product)
	if result.Error != nil {
		return fmt.Errorf("error adding product: %v", result.Error)
	}
	return nil
}

func GetProducts(db *gorm.DB) ([]Product, error) {
	var product []Product
	result := db.Find(&product)

	if result.Error != nil {
		return nil, fmt.Errorf("error fetching products: %v", result.Error)
	}
	return product, nil
}

func UpdateProduct(db *gorm.DB, id string, updatedProduct Product) (*Product, error) {
	var product Product
	if err := db.First(&product, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("product not found")
	}

	log.Println(updatedProduct)
	product.Name = updatedProduct.Name
	product.Description = updatedProduct.Description
	product.Quantity = updatedProduct.Quantity
	product.Price = updatedProduct.Price

	if err := db.Save(&product).Error; err != nil {
		return nil, fmt.Errorf("failed to update product")
	}

	return &product, nil
}

func DeleteProduct(db *gorm.DB, id string) error {
	var product Product
	if err := db.First(&product, "id = ?", id).Error; err != nil {
		return fmt.Errorf("product not found")
	}

	if err := db.Delete(&product).Error; err != nil {
		return fmt.Errorf("failed to delete product")
	}

	return nil
}

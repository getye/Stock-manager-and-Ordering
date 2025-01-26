package models

import (
	"fmt"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name     string    `gorm:"size:255;not null" json:"name"`
	Email    string    `gorm:"size:255;unique;not null" json:"email"`
	Phone    string    `gorm:"size:255;not null" json:"phone"`
	Password string    `gorm:"size:255;not null" json:"password"`
	Role     string    `gorm:"size:255;not null" json:"role"`
	Status   string    `gorm:"size:255;not null" json:"status"`
}

func InsertUser(db *gorm.DB, user *User) error {

	result := db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("error inserting user: %v", result.Error)
	}
	return nil
}

func GetUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Select("id, name, role, email, phone, status").Find(&users)

	if result.Error != nil {
		return nil, fmt.Errorf("error fetching users: %v", result.Error)
	}
	return users, nil
}

func UpdateUser(db *gorm.DB, id string, updatedUser User) (*User, error) {
	var user User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Role = updatedUser.Role
	user.Phone = updatedUser.Phone

	if err := db.Save(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to update user")
	}

	return &user, nil
}

func DeleteUser(db *gorm.DB, id string) error {
	var user User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return fmt.Errorf("user not found")
	}

	if err := db.Delete(&user).Error; err != nil {
		return fmt.Errorf("failed to delete user")
	}

	return nil
}

func UpdateUserStatus(db *gorm.DB, id string, userStatus User) (*User, error) {
	var user User
	if err := db.First(&user, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}

	user.Status = userStatus.Status

	if err := db.Save(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to update user status")
	}

	return &user, nil
}

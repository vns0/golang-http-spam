package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/database"
)

func GetUsers(c *gin.Context) {
	var users []database.User
	database.Db.Unscoped().Find(&users)

	c.JSON(200, gin.H{"users": users})
}

func CreateUser(c *gin.Context) {
	var newUser database.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var existingUser database.User
	if err := database.Db.Where("user_id = ?", newUser.UserID).Unscoped().First(&existingUser).Error; err == nil && existingUser.DeletedAt.Valid {
		existingUser.DeletedAt = gorm.DeletedAt{}
		database.Db.Save(&existingUser)
		c.JSON(200, gin.H{"user": existingUser, "message": "User restored successfully"})
		return
	}
	database.Db.Create(&newUser)
	c.JSON(200, gin.H{"user": newUser, "message": "User created successfully"})
}

func DeleteUser(c *gin.Context) {
	var newUser database.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	existingUser := database.User{}
	result := database.Db.Where("user_id = ?", newUser.UserID).First(&existingUser)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	result = database.Db.Delete(&existingUser)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(200, gin.H{"user": existingUser, "message": "User delete successfully"})
}

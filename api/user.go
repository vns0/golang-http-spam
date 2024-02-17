package api

import (
	"github.com/gin-gonic/gin"
	"main.go/database"
)

func GetUsers(c *gin.Context) {
	var users []database.User
	database.Db.Find(&users)

	c.JSON(200, gin.H{"users": users})
}

func CreateUser(c *gin.Context) {
	var newUser database.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.Db.Create(&newUser)
	c.JSON(200, gin.H{"user": newUser})
}

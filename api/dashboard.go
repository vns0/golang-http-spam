package api

import (
	"github.com/gin-gonic/gin"
	"main.go/database"
)

func GetStats(c *gin.Context) {
	var countUsers int64
	database.Db.Find(&database.User{}).Count(&countUsers)
	var countAttack int64
	database.Db.Find(&database.AttackHistory{}).Count(&countAttack)
	c.JSON(200, gin.H{"countAttack": countAttack, "countUsers": countUsers})
}

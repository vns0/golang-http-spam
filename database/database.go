package database

import (
	"errors"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var AdminUserID int64 = 696300339

var db *gorm.DB

type AttackHistory struct {
	gorm.Model
	UserID    *int64
	Command   string
	AttackURL string
	Success   bool
}

type Admin struct {
	gorm.Model
	UserID int64 `gorm:"uniqueIndex"`
}

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("spam_bot.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Admin{}, &AttackHistory{})
}

func SaveAttackHistory(userID *int64, command, attackURL string, success bool) {
	db.Create(&AttackHistory{
		UserID:    userID,
		Command:   command,
		AttackURL: attackURL,
		Success:   success,
	})
}

func IsAdmin(userID int64) bool {
	var admin Admin
	if err := db.Where("user_id = ?", userID).First(&admin).Error; err != nil {
		return false
	}
	return true
}

func AddAdmin(userID int64) {
	var admin Admin
	result := db.Where("user_id = ?", userID).First(&admin)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(&Admin{UserID: userID})
		fmt.Println("Admin added:", userID)
	}
}

func DeleteAdmin(userID int64) {
	db.Where("user_id=?", userID).Delete(&Admin{})
}

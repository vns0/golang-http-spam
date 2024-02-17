package database

import (
	"errors"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"time"
)

var AdminUserID int64 = 696300339

var Db *gorm.DB

type AttackHistory struct {
	gorm.Model
	UserID    *int64
	Command   string
	AttackURL string
	Success   bool
}

type User struct {
	gorm.Model
	UserID int64 `gorm:"uniqueIndex"`
}

type AuthCode struct {
	gorm.Model
	TelegramID int64
	Code       string
	IsUsed     bool
	CreatedAt  time.Time
}

func InitDB() {
	var err error
	Db, err = gorm.Open(sqlite.Open("spam_bot.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&User{}, &AttackHistory{}, &AuthCode{})
}

func SaveAttackHistory(userID *int64, command, attackURL string, success bool) {
	Db.Create(&AttackHistory{
		UserID:    userID,
		Command:   command,
		AttackURL: attackURL,
		Success:   success,
	})
}

func IsAdmin(userID int64) bool {
	var admin User
	if err := Db.Where("user_id = ?", userID).First(&admin).Error; err != nil {
		return false
	}
	return true
}

func AddAdmin(userID int64) {
	var admin User
	result := Db.Where("user_id = ?", userID).First(&admin)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		Db.Create(&User{UserID: userID})
		fmt.Println("Admin added:", userID)
	}
}

func DeleteAdmin(userID int64) {
	Db.Where("user_id=?", userID).Delete(&User{})
}

func GetAuthCode(code string) (*AuthCode, error) {
	var authCode AuthCode
	err := Db.Where("code = ? AND is_used = ?", code, false).First(&authCode).Error
	if err != nil {
		return nil, err
	}
	return &authCode, nil
}

func DeleteAuthCode(code string) bool {
	if err := Db.Where("code=?", code).Delete(&AuthCode{}).Error; err != nil {
		return false
	}
	return true
}

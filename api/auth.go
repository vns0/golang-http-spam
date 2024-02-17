package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"main.go/database"
	"math/rand"
	"sync"
	"time"
)

// @TODO: update settings
const (
	AccessTokenSecret  = "qwerty"
	RefreshTokenSecret = "qwerty"
	TokenDuration      = time.Minute * 1440
)

var authCodes = make(map[int64]string)
var mu sync.Mutex

func generateAuthCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func ProcessAuthCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	userID := msg.From.ID
	code := generateAuthCode()
	newAuthCode := database.AuthCode{
		TelegramID: int64(userID),
		Code:       code,
		IsUsed:     false,
		CreatedAt:  time.Now(),
	}

	if err := database.Db.Create(&newAuthCode).Error; err != nil {
		authMsg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Failed to create auth code"))
		_, err := bot.Send(authMsg)
		if err != nil {
			log.Printf("Failed to send authorization code to %d: %v", chatID, err)
			return
		}
		return
	}
	authMsg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Your authorization code: %s", code))
	_, err := bot.Send(authMsg)
	if err != nil {
		log.Printf("Failed to send authorization code to %d: %v", chatID, err)
		return
	}
	mu.Lock()
	authCodes[int64(userID)] = code
	mu.Unlock()
}

/*
*
@description - 5 min validation auth code
*/
func isAuthCodeValid(authCode database.AuthCode) bool {
	return time.Since(authCode.CreatedAt) <= 5*time.Minute
}

func GenerateTokens(userID int64) (string, string, error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	accessTokenClaims["authorized"] = true
	accessTokenClaims["user_id"] = userID
	accessTokenClaims["exp"] = time.Now().Add(TokenDuration).Unix()
	accessTokenString, err := accessToken.SignedString([]byte(AccessTokenSecret))
	if err != nil {
		return "", "", err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["user_id"] = userID
	refreshTokenClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	refreshTokenString, err := refreshToken.SignedString([]byte(RefreshTokenSecret))
	if err != nil {
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}

func AuthLogin(c *gin.Context) {
	var requestBody struct {
		Code string `json:"code" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	code := requestBody.Code
	if code == "" {
		c.JSON(400, gin.H{"error": "Code parameter is missing"})
		return
	}
	authCode, err := database.GetAuthCode(code)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid or expired code"})
		return
	}
	if !isAuthCodeValid(*authCode) {
		database.DeleteAuthCode(code)
		c.JSON(400, gin.H{"error": "Invalid or expired code"})
		return
	}
	authCode.IsUsed = true
	if err := database.Db.Save(authCode).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update auth code status"})
		return
	}
	accessToken, refreshToken, err := GenerateTokens(authCode.TelegramID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate tokens"})
		return
	}
	c.JSON(200, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}

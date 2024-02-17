package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main.go/api"
	"net/http"
	"strings"
)

type AccessTokenClaims struct {
	jwt.StandardClaims
	TelegramID int64 `json:"telegramID"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := verifyAccessToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("telegramID", accessToken.TelegramID)
		c.Next()
	}
}

func verifyAccessToken(c *gin.Context) (*AccessTokenClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("missing Authorization header")
	}
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return nil, fmt.Errorf("invalid Authorization header format")
	}
	accessToken := splitToken[1]
	token, err := jwt.ParseWithClaims(accessToken, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(api.AccessTokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid access token")
	}
	return claims, nil
}

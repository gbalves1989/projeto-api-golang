package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateJWT(userID uint, secretKey string) string {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		Log(ERROR, "Erro ao assinar token JWT.")
	}

	return tokenString
}

func ValidateJWT(tokenString string, secretKey string) uint {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		Log(WARN, "Erro ao validar token JWT.")
		return 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		return userID
	}

	return 0
}

func GetUserIDFromContext(c *gin.Context) uint {
	userID, exists := c.Get("userID")
	if !exists {
		Log(ERROR, "User ID não encontrado no contexto Gin.")
		return 0
	}

	id, ok := userID.(uint)
	if !ok {
		Log(ERROR, "Tipo de User ID inválido no contexto Gin.")
		return 0
	}

	return id
}

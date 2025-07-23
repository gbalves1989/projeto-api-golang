package middlewares

import (
	"net/http"
	"strings"

	"github.com/gbalves1989/projeto-api-golang/config"
	"github.com/gbalves1989/projeto-api-golang/constants"
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Code:       http.StatusBadRequest,
				StatusCode: constants.UNATHORIZED,
				Message:    "Token de autenticação inválido.",
			})

			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Code:       http.StatusUnauthorized,
				StatusCode: constants.UNATHORIZED,
				Message:    "Formato do token inválido.",
			})

			return
		}

		tokenString := parts[1]
		userID := utils.ValidateJWT(tokenString, config.AppConfig.JWTSecret)
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
				Code:       http.StatusUnauthorized,
				StatusCode: constants.UNATHORIZED,
				Message:    "Token inválido.",
			})

			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

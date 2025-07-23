package controllers

import (
	services "github.com/gbalves1989/projeto-api-golang/http/services/auth"
	"github.com/gin-gonic/gin"
)

func AuthRegisterController(c *gin.Context) {
	services.AuthRegisterService(c)
}

func AuthLoginController(c *gin.Context) {
	services.AuthLoginService(c)
}

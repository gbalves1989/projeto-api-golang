package services

import (
	"net/http"

	"github.com/gbalves1989/projeto-api-golang/config"
	"github.com/gbalves1989/projeto-api-golang/constants"
	"github.com/gbalves1989/projeto-api-golang/database/repositories"
	"github.com/gbalves1989/projeto-api-golang/http/requests"
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"
)

func AuthLoginService(c *gin.Context) {
	var input requests.AuthLoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Dados de entrada inválido.",
		})

		return
	}

	user := repositories.UserShowByEmailRepository(input.Email)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Code:       http.StatusNotFound,
			StatusCode: constants.NOT_FOUND,
			Message:    "E-mail não encontrado.",
		})

		return
	}

	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{
			Code:       http.StatusUnauthorized,
			StatusCode: constants.UNATHORIZED,
			Message:    "Credenciais inválidas.",
		})

		return
	}

	token := utils.GenerateJWT(user.ID, config.AppConfig.JWTSecret)
	utils.Log(utils.INFO, "Login bem-sucedido.")
	c.JSON(http.StatusOK, responses.TokenResponse{
		Code:        http.StatusOK,
		StatusCode:  constants.OK,
		Message:     "Usuário logado com sucesso.",
		AccessToken: token,
	})
}

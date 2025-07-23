package services

import (
	"net/http"

	"github.com/gbalves1989/projeto-api-golang/constants"
	"github.com/gbalves1989/projeto-api-golang/database/repositories"
	"github.com/gbalves1989/projeto-api-golang/http/requests"
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"
)

func UserUpdatePasswordService(c *gin.Context) {
	var input requests.UserUpdatePasswordRequest
	var hash string

	id := utils.GetUserIDFromContext(c)
	user := repositories.UserShowByIDRepository(id)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Dados de entrada inválido.",
		})

		return
	}

	if input.Password != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Senhas não conferem.",
		})

		return
	}

	hash = utils.HashPassword(input.Password)
	user = repositories.UserUpdatePasswordRepository(hash, user)
	utils.Log(utils.INFO, "Senha atualizada com sucesso.")
	c.JSON(http.StatusAccepted, responses.UserResponse{
		Code:       http.StatusAccepted,
		StatusCode: constants.ACCEPTED,
		Message:    "Senha atualizada com sucesso.",
		Data:       user.ToUserResponse(),
	})
}

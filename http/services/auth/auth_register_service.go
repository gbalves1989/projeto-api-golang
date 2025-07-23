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

func AuthRegisterService(c *gin.Context) {
	var input requests.AuthRegisterRequest

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

	existsEmail := repositories.UserShowByEmailRepository(input.Email)
	if existsEmail.ID > 0 {
		c.JSON(http.StatusConflict, responses.ErrorResponse{
			Code:       http.StatusConflict,
			StatusCode: constants.BAD_REQUEST,
			Message:    "E-mail já está cadastrado.",
		})

		return
	}

	hashedPassword := utils.HashPassword(input.Password)
	cretatedUser := repositories.UserStoreRepository(input, hashedPassword)

	utils.Log(utils.INFO, "Usuário cadastrado com sucesso.")
	c.JSON(http.StatusCreated, responses.UserResponse{
		Code:       http.StatusCreated,
		StatusCode: constants.CREATED,
		Message:    "Usuário cadastrado com sucesso.",
		Data:       cretatedUser.ToUserResponse(),
	})
}

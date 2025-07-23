package services

import (
	"net/http"

	"github.com/gbalves1989/projeto-api-golang/constants"
	"github.com/gbalves1989/projeto-api-golang/database/repositories"
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"
)

func UserMeService(c *gin.Context) {
	id := utils.GetUserIDFromContext(c)
	user := repositories.UserShowByIDRepository(id)
	c.JSON(http.StatusOK, responses.UserResponse{
		Code:       http.StatusOK,
		StatusCode: constants.OK,
		Message:    "Usuário encontrado com sucesso.",
		Data:       user.ToUserResponse(),
	})
}

package services

import (
	"net/http"
	"os"
	"strings"

	"github.com/gbalves1989/projeto-api-golang/config"
	"github.com/gbalves1989/projeto-api-golang/constants"
	"github.com/gbalves1989/projeto-api-golang/database/repositories"
	"github.com/gbalves1989/projeto-api-golang/http/requests"
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"
)

func UserDestroyService(c *gin.Context) {
	var input requests.UserDeleteRequest

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

	if user.AvatarURL != "" {
		filePath := strings.Replace(user.AvatarURL, "/uploads", config.AppConfig.UploadDir, 1)
		if _, err := os.Stat(filePath); err == nil {
			if err := os.Remove(filePath); err != nil {
				utils.Log(utils.WARN, "Falha ao deletar arquivo de avatar para usuário %d: %v", user.ID, err)
			} else {
				utils.Log(utils.INFO, "Arquivo de avatar deletado para usuário %d", user.ID)
			}
		}
	}

	repositories.UserDestroyRepository(user)
	utils.Log(utils.INFO, "Usuário deletado com sucesso.")
	c.Status(http.StatusNoContent)
}

package services

import (
	"net/http"
	"os"
	"strings"

	"github.com/gbalves1989/projeto-api-golang/config"
	"github.com/gbalves1989/projeto-api-golang/constants"
	"github.com/gbalves1989/projeto-api-golang/database/repositories"
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"
)

func UserUploadService(c *gin.Context) {
	id := utils.GetUserIDFromContext(c)
	user := repositories.UserShowByIDRepository(id)

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Nenhum arquivo 'avatar' fornecido",
		})

		return
	}

	validateFile := utils.ValidateImageFile(file)
	if !validateFile {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Formato de arquivo inválido",
		})

		return
	}

	if user.AvatarURL != "" {
		oldFilePath := strings.Replace(user.AvatarURL, "/uploads", config.AppConfig.UploadDir, 1)
		if _, err := os.Stat(oldFilePath); err == nil {
			if err := os.Remove(oldFilePath); err != nil {
				utils.Log(utils.WARN, "Falha ao deletar avatar antigo para usuário %d: %v", user.ID, err)
			}
		}
	}

	filePath := utils.SaveUploadedFile(file, config.AppConfig.UploadDir)
	if filePath == "" {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Code:       http.StatusInternalServerError,
			StatusCode: constants.INTERNAL_SERVER_ERROR,
			Message:    "Erro ao tentar salvar o arquivo",
		})

		return
	}

	user.AvatarURL = filePath
	user = repositories.UserUploadRepository(filePath, user)
	utils.Log(utils.INFO, "Avatar atualizado com sucesso.")
	c.JSON(http.StatusAccepted, responses.UserResponse{
		Code:       http.StatusAccepted,
		StatusCode: constants.ACCEPTED,
		Message:    "Avatar atualizado com sucesso.",
		Data:       user.ToUserResponse(),
	})
}

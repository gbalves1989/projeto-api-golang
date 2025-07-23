package services

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gbalves1989/projeto-api-golang/config"
	"github.com/gbalves1989/projeto-api-golang/constants"
	"github.com/gbalves1989/projeto-api-golang/database/repositories"
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"
)

func ProductUploadService(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "ID do produto inválido.",
		})

		return
	}

	product := repositories.ProductShowRepository(productID)
	if product.ID == 0 {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Code:       http.StatusNotFound,
			StatusCode: constants.NOT_FOUND,
			Message:    "Produto não encontrado.",
		})

		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Nenhum arquivo 'image' fornecido",
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

	if product.ImageURL != "" {
		oldFilePath := strings.Replace(product.ImageURL, "/uploads", config.AppConfig.UploadDir, 1)
		if _, err := os.Stat(oldFilePath); err == nil {
			if err := os.Remove(oldFilePath); err != nil {
				utils.Log(utils.WARN, "Falha ao tentar deletar imagem do produto %d: %v", product.ID, err)
			}
		}
	}

	filePath := utils.SaveUploadedFile(file, config.AppConfig.UploadDir)
	if filePath == "" {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Code:       http.StatusInternalServerError,
			StatusCode: constants.INTERNAL_SERVER_ERROR,
			Message:    "Erro ao tentar salvar o arquivo.",
		})

		return
	}

	product = repositories.ProductUploadRepository(filePath, product)
	product = repositories.CategoryByProductRepository(product)
	utils.Log(utils.INFO, "Imagem do produto atualizada com sucesso.")
	c.JSON(http.StatusAccepted, responses.ProductResponse{
		Code:       http.StatusAccepted,
		StatusCode: constants.ACCEPTED,
		Message:    "Imagem do produto atualizada com sucesso.",
		Data:       product.ToProductResponse(),
	})
}

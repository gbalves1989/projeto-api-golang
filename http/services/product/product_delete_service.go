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

func ProductDestroyService(c *gin.Context) {
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

	if product.ImageURL != "" {
		filePath := strings.Replace(product.ImageURL, "/uploads", config.AppConfig.UploadDir, 1)
		if _, err := os.Stat(filePath); err == nil {
			if err := os.Remove(filePath); err != nil {
				utils.Log(utils.WARN, "Falha ao deletar arquivo de imagem do produto %d: %v", product.ID, err)
			} else {
				utils.Log(utils.INFO, "Arquivo de imagem deletado para o produto %d", product.ID)
			}
		}
	}

	repositories.ProductDestroyRepository(product)
	utils.Log(utils.INFO, "Produto deletado com sucesso.")
	c.Status(http.StatusNoContent)
}

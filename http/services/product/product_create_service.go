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

func ProductStoreService(c *gin.Context) {
	var input requests.ProductCreateOrUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Dados de entrada inválido.",
		})

		return
	}

	existsName := repositories.ProductShowByNameRepository(input.Name)
	if existsName.ID > 0 {
		c.JSON(http.StatusConflict, responses.ErrorResponse{
			Code:       http.StatusConflict,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Nome do produto já está cadastrado.",
		})

		return
	}

	category := repositories.CategoryShowRepository(uint64(input.CategoryID))
	if category.ID == 0 {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Code:       http.StatusNotFound,
			StatusCode: constants.NOT_FOUND,
			Message:    "Categoria não encontrada.",
		})

		return
	}

	createdProduct := repositories.ProductStoreRepository(input)
	createdProduct = repositories.CategoryByProductRepository(createdProduct)
	utils.Log(utils.INFO, "Produto cadastrado com sucesso.")
	c.JSON(http.StatusCreated, responses.ProductResponse{
		Code:       http.StatusCreated,
		StatusCode: constants.CREATED,
		Message:    "Produto cadastrado com sucesso.",
		Data:       createdProduct.ToProductResponse(),
	})
}

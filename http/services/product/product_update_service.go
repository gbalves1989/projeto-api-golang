package services

import (
	"net/http"
	"strconv"

	"github.com/gbalves1989/projeto-api-golang/constants"
	"github.com/gbalves1989/projeto-api-golang/database/repositories"
	"github.com/gbalves1989/projeto-api-golang/http/requests"
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"
)

func ProductUpdateService(c *gin.Context) {
	var input requests.ProductCreateOrUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Dados de entrada inválido.",
		})

		return
	}

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

	existsID := repositories.ProductShowRepository(productID)
	if existsID.ID == 0 {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Code:       http.StatusNotFound,
			StatusCode: constants.NOT_FOUND,
			Message:    "Produto não encontrado.",
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

	updatedProduct := repositories.ProductUpdateRepository(input, existsID)
	updatedProduct = repositories.CategoryByProductRepository(updatedProduct)
	utils.Log(utils.INFO, "Produto atualizado com sucesso.")
	c.JSON(http.StatusAccepted, responses.ProductResponse{
		Code:       http.StatusAccepted,
		StatusCode: constants.ACCEPTED,
		Message:    "Produto atualizado com sucesso.",
		Data:       updatedProduct.ToProductResponse(),
	})
}

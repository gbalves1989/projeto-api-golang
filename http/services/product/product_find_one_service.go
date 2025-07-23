package services

import (
	"net/http"
	"strconv"

	"github.com/gbalves1989/projeto-api-golang/constants"
	"github.com/gbalves1989/projeto-api-golang/database/repositories"
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"
)

func ProductShowService(c *gin.Context) {
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

	product = repositories.CategoryByProductRepository(product)
	utils.Log(utils.INFO, "Produto encontrado com sucesso.")
	c.JSON(http.StatusOK, responses.ProductResponse{
		Code:       http.StatusOK,
		StatusCode: constants.OK,
		Message:    "Produto encontrado com sucesso.",
		Data:       product.ToProductResponse(),
	})
}

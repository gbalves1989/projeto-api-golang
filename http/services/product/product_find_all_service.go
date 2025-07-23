package services

import (
	"net/http"

	"github.com/gbalves1989/projeto-api-golang/constants"
	"github.com/gbalves1989/projeto-api-golang/database/models"
	"github.com/gbalves1989/projeto-api-golang/database/repositories"
	"github.com/gbalves1989/projeto-api-golang/http/requests"
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"
)

func ProductIndexService(c *gin.Context) {
	var products []models.Product
	var pagination requests.PaginationRequest

	c.ShouldBindQuery(&pagination)
	total := repositories.CountProductsRepository()
	products = repositories.ProductIndexRepository(pagination)

	hasNext := int64(pagination.Offset)+int64(len(products)) < total
	hasPrevious := pagination.Offset > 0

	productResponses := make([]responses.ProductData, len(products))
	for i, product := range products {
		productResponses[i] = product.ToProductResponse()
	}

	response := responses.PaginationResponse{
		Total:       total,
		Limit:       pagination.Limit,
		Offset:      pagination.Offset,
		Data:        productResponses,
		HasNext:     hasNext,
		HasPrevious: hasPrevious,
	}

	utils.Log(utils.INFO, "Produtos encontrados com sucesso.")
	c.JSON(http.StatusOK, responses.ProductResponse{
		Code:       http.StatusOK,
		StatusCode: constants.OK,
		Message:    "Produtos encontrados com sucesso.",
		Data:       response,
	})
}

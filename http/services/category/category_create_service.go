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

func CategoryStoreService(c *gin.Context) {
	var input requests.CategoryCreateOrUpdateRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Dados de entrada inválido.",
		})

		return
	}

	existsName := repositories.CategoryShowByNameRepository(input.Name)
	if existsName.ID > 0 {
		c.JSON(http.StatusConflict, responses.ErrorResponse{
			Code:       http.StatusConflict,
			StatusCode: constants.CONFLICT,
			Message:    "Nome da categoria já está cadastrada.",
		})

		return
	}

	createdCategory := repositories.CategoryStoreRepository(input)
	createdCategory = repositories.ProductsByCategoryRepository(createdCategory, uint64(createdCategory.ID))
	productsList := make([]responses.ProductsByCategoryData, len(createdCategory.Products))
	for i, p := range createdCategory.Products {
		productsList[i] = responses.ProductsByCategoryData{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			ImageURL:    p.ImageURL,
		}
	}

	utils.Log(utils.INFO, "Categoria criada com sucesso.")
	c.JSON(http.StatusCreated, responses.CategoryResponse{
		Code:       http.StatusCreated,
		StatusCode: constants.CREATED,
		Message:    "Categoria cadastrado com sucesso.",
		Data: responses.CategoryData{
			ID:          createdCategory.ID,
			Name:        createdCategory.Name,
			Description: createdCategory.Description,
			Products:    productsList,
		},
	})
}

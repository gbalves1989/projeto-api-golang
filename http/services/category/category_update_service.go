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

func CategoryUpdateService(c *gin.Context) {
	var input requests.CategoryCreateOrUpdateRequest

	id := c.Param("id")
	categoryID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "ID de categoria inválido.",
		})

		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Code:       http.StatusBadRequest,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Dados de entrada inválido.",
		})

		return
	}

	category := repositories.CategoryShowRepository(categoryID)
	if category.ID == 0 {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Code:       http.StatusNotFound,
			StatusCode: constants.NOT_FOUND,
			Message:    "Categoria não encontrada.",
		})

		return
	}

	existsName := repositories.CategoryShowByNameRepository(input.Name)
	if existsName.ID > 0 {
		c.JSON(http.StatusConflict, responses.ErrorResponse{
			Code:       http.StatusConflict,
			StatusCode: constants.BAD_REQUEST,
			Message:    "Nome da categoria já está cadastrada.",
		})

		return
	}

	category = repositories.CategoryUpdateRepository(input, category)
	category = repositories.ProductsByCategoryRepository(category, categoryID)
	productsList := make([]responses.ProductsByCategoryData, len(category.Products))
	for i, p := range category.Products {
		productsList[i] = responses.ProductsByCategoryData{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			ImageURL:    p.ImageURL,
		}
	}

	utils.Log(utils.INFO, "Categoria atualizada com sucesso.")
	c.JSON(http.StatusAccepted, responses.CategoryResponse{
		Code:       http.StatusAccepted,
		StatusCode: constants.ACCEPTED,
		Message:    "Usuário atualizado com sucesso.",
		Data: responses.CategoryData{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			Products:    productsList,
		},
	})
}

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

func CategoryShowService(c *gin.Context) {
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

	category := repositories.CategoryShowRepository(categoryID)
	if category.ID == 0 {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{
			Code:       http.StatusNotFound,
			StatusCode: constants.NOT_FOUND,
			Message:    "Categoria não encontrada.",
		})

		return
	}

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

	utils.Log(utils.INFO, "Categoria encontrada com sucesso")
	c.JSON(http.StatusOK, responses.CategoryResponse{
		Code:       http.StatusOK,
		StatusCode: constants.OK,
		Message:    "Categoria encontrada com sucesso.",
		Data: responses.CategoryData{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			Products:    productsList,
		},
	})
}

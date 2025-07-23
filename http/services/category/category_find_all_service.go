package services

import (
	"net/http"

	"github.com/gbalves1989/projeto-api-golang/constants"
	"github.com/gbalves1989/projeto-api-golang/database/models"
	"github.com/gbalves1989/projeto-api-golang/database/repositories"
	"github.com/gbalves1989/projeto-api-golang/http/requests"
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"github.com/gin-gonic/gin"
)

func CategoryIndexService(c *gin.Context) {
	var categories []models.Category
	var pagination requests.PaginationRequest

	c.ShouldBindQuery(&pagination)
	total := repositories.CountCategoriesRepository()
	categories = repositories.CategoryIndexRepository(pagination)

	hasNext := int64(pagination.Offset)+int64(len(categories)) < total
	hasPrevious := pagination.Offset > 0

	categoriesList := make([]responses.CategoryListData, len(categories))
	for i, p := range categories {
		categoriesList[i] = responses.CategoryListData{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
		}
	}

	response := responses.PaginationResponse{
		Total:       total,
		Limit:       pagination.Limit,
		Offset:      pagination.Offset,
		Data:        categoriesList,
		HasNext:     hasNext,
		HasPrevious: hasPrevious,
	}

	c.JSON(http.StatusOK, responses.CategoryListResponse{
		Code:       http.StatusOK,
		StatusCode: constants.OK,
		Message:    "Categorias encontradas com sucesso.",
		Data:       response,
	})
}

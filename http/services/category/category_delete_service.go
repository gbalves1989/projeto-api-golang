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

func CategoryDestroyService(c *gin.Context) {
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
	if len(category.Products) > 0 {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse{
			Code:       http.StatusInternalServerError,
			StatusCode: constants.INTERNAL_SERVER_ERROR,
			Message:    "Categoria possui produtos cadastrados.",
		})

		return
	}

	repositories.CategoryDestroyRepository(category)
	utils.Log(utils.INFO, "Categoria deletada com sucesso.")
	c.Status(http.StatusNoContent)
}

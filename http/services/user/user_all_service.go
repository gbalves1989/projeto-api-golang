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

func UserIndexService(c *gin.Context) {
	var users []models.User
	var pagination requests.PaginationRequest

	c.ShouldBindQuery(&pagination)
	total := repositories.CountUsersRepository()
	users = repositories.UserIndexRepository(pagination)

	hasNext := int64(pagination.Offset)+int64(len(users)) < total
	hasPrevious := pagination.Offset > 0

	response := responses.PaginationResponse{
		Total:       total,
		Limit:       pagination.Limit,
		Offset:      pagination.Offset,
		Data:        users,
		HasNext:     hasNext,
		HasPrevious: hasPrevious,
	}

	c.JSON(http.StatusOK, responses.UserListResponse{
		Code:       http.StatusOK,
		StatusCode: constants.OK,
		Message:    "Usuários encontrados com sucesso.",
		Data:       response,
	})
}

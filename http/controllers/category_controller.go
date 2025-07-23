package controllers

import (
	services "github.com/gbalves1989/projeto-api-golang/http/services/category"
	"github.com/gin-gonic/gin"
)

// CreateCategory godoc
// @Summary Cadastrar uma nova categoria
// @Description Responsável por cadastrar uma nova categoria
// @Tags Categorias
// @Accept json
// @Produce json
// @Param category body requests.CategoryCreateOrUpdateRequest true "Dados da categoria"
// @Success 201 {object} responses.CategoryResponse "Retorna os dados da categoria"
// @Failure 400 {object} responses.ErrorResponse "Dados de entrada inválido"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 409 {object} responses.ErrorResponse "Nome da categoria já existe"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Security ApiKeyAuth
// @Router /categories [post]
func CategoryStoreController(c *gin.Context) {
	services.CategoryStoreService(c)
}

// ShowCategory godoc
// @Summary Retornar uma categoria por ID
// @Description Responsável por retornar uma categoria por ID
// @Tags Categorias
// @Accept json
// @Produce json
// @Param id path int true "ID da Categoria"
// @Success 200 {object} responses.CategoryResponse "Retorna uma categoria"
// @Failure 400 {object} responses.ErrorResponse "ID de categoria inválido"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 404 {object} responses.ErrorResponse "Categoria não encontrada"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Security ApiKeyAuth
// @Router /categories/{id} [get]
func CategoryShowController(c *gin.Context) {
	services.CategoryShowService(c)
}

// IndexCategory godoc
// @Summary Retonar uma lista categorias
// @Description Responsável por retornar uma lista categorias
// @Tags Categorias
// @Accept json
// @Produce json
// @Success 200 {object} responses.CategoryListResponse "Retorna uma lista de categorias"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Security ApiKeyAuth
// @Router /categories [get]
func CategoryIndexController(c *gin.Context) {
	services.CategoryIndexService(c)
}

// UpdateCategory godoc
// @Summary Atualizar uma categoria por ID
// @Description Responsável por atualizar as informações de uma categoria por ID
// @Tags Categorias
// @Accept json
// @Produce json
// @Param id path int true "ID da Categoria"
// @Param category body requests.CategoryCreateOrUpdateRequest true "Dados da categoria"
// @Success 202 {object} responses.CategoryResponse "Retorna uma categoria atualizada"
// @Failure 400 {object} responses.ErrorResponse "ID de categoria inválido | Dados de entrada inválido"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 404 {object} responses.ErrorResponse "Categoria não encontrada"
// @Failure 409 {object} responses.ErrorResponse "Nome da categoria já está cadastrada",
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Security ApiKeyAuth
// @Router /categories/{id} [put]
func CategoryUpdateController(c *gin.Context) {
	services.CategoryUpdateService(c)
}

// DestroyCategory godoc
// @Summary Remover uma categoria por ID
// @Description Responsável por remover uma categoria por ID
// @Tags Categorias
// @Accept json
// @Produce json
// @Param id path int true "ID da Categoria"
// @Success 204 "No Content"
// @Failure 400 {object} responses.ErrorResponse "ID de categoria inválido"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 404 {object} responses.ErrorResponse "Categoria não encontrada"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno | Categoria possui produtos cadastrados"
// @Security ApiKeyAuth
// @Router /categories/{id} [delete]
func CategoryDestroyController(c *gin.Context) {
	services.CategoryDestroyService(c)
}

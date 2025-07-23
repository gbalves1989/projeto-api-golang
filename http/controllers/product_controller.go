package controllers

import (
	services "github.com/gbalves1989/projeto-api-golang/http/services/product"
	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Cadastrar um novo produto
// @Description Responsável por cadastrar um novo produto
// @Tags Produtos
// @Accept json
// @Produce json
// @Param product body requests.ProductCreateOrUpdateRequest true "Dados do produto"
// @Success 201 {object} responses.ProductResponse "Retorna os dados do produto"
// @Failure 400 {object} responses.ErrorResponse "Dados de entrada inválido"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 409 {object} responses.ErrorResponse "Nome do produto já existe"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Security ApiKeyAuth
// @Router /products [post]
func ProductStoreController(c *gin.Context) {
	services.ProductStoreService(c)
}

// IndexProduct godoc
// @Summary Listar produtos
// @Description Responsável por retornar uma lista de produtos
// @Tags Produtos
// @Accept json
// @Produce json
// @Success 200 {object} responses.ProductResponse "Retorna uma lista de produtos"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Security ApiKeyAuth
// @Router /products [get]
func ProductIndexController(c *gin.Context) {
	services.ProductIndexService(c)
}

// ShowProduct godoc
// @Summary Retorna um produto por ID
// @Description Responsável por retornar um produto por ID
// @Tags Produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 200 {object} responses.ProductResponse "Retorna um produto"
// @Failure 400 {object} responses.ErrorResponse "ID do produto inválido"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 404 {object} responses.ErrorResponse "Produto não encontrado"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Security ApiKeyAuth
// @Router /products/{id} [get]
func ProductShowController(c *gin.Context) {
	services.ProductShowService(c)
}

// UpdateProduct godoc
// @Summary Atualizar um produto por ID
// @Description Responsável por atualizar as informações de um produto por ID
// @Tags Produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Param product body requests.ProductCreateOrUpdateRequest true "Dados do produto"
// @Success 202 {object} responses.ProductResponse "Retorna um produto atualizado"
// @Failure 400 {object} responses.ErrorResponse "Dados de entrada inválido | ID do produto inválido"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 404 {object} responses.ErrorResponse "Produto não encontrado | Categoria não encontrada"
// @Failure 409 {object} responses.ErrorResponse "Nome do produto já está cadastrado"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Security ApiKeyAuth
// @Router /categories/{id} [put]
func ProductUpdateController(c *gin.Context) {
	services.ProductUpdateService(c)
}

// UploadProduct godoc
// @Summary Faz upload da imagem deum produto por ID
// @Description Responsável por atualizar a imagem de um produto por ID
// @Tags Produtos
// @Security ApiKeyAuth
// @Accept mpfd
// @Produce json
// @Param id path int true "ID do produto"
// @Param image formData file true "Arquivo de imagem do produto (JPEG, PNG, GIF, máx 5MB)"
// @Success 202 {object} responses.ProductResponse "Retorna um produto com a imagem atualizada"
// @Failure 400 {object} responses.ErrorResponse "ID do produto inválido | Nenhum arquivo 'image' fornecido | Formato de arquivo inválido"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 404 {object} responses.ErrorResponse "Produto não encontrado"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro ao tentar salvar o arquivo"
// @Router /products/upload/{id} [patch]
func ProductUploadController(c *gin.Context) {
	services.ProductUploadService(c)
}

// DestroyProduct godoc
// @Summary Remove um produto por ID
// @Description Responsável por remover um produto por ID
// @Tags Produtos
// @Accept json
// @Produce json
// @Param id path int true "ID do produto"
// @Success 204 "No Content"
// @Failure 400 {object} responses.ErrorResponse "ID do produto inválido"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 404 {object} responses.ErrorResponse "Produto não encontrado"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Security ApiKeyAuth
// @Router /products/{id} [delete]
func ProductDestroyController(c *gin.Context) {
	services.ProductDestroyService(c)
}

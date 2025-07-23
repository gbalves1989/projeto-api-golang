package controllers

import (
	services "github.com/gbalves1989/projeto-api-golang/http/services/user"
	"github.com/gin-gonic/gin"
)

// MeUser godoc
// @Summary Retornar informações de um usuário logado
// @Description Responável por retornar as informações de um usuário logado
// @Tags Usuários
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} responses.UserResponse "Retorna informações do usuário"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Router /users/me [get]
func UserMeController(c *gin.Context) {
	services.UserMeService(c)
}

// UpdateUser godoc
// @Summary Atualizar as informações de um usuário
// @Description Responsável por atualizar as informações de um usuário
// @Tags Usuários
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body requests.UserUpdateRequest true "Dados do usuário"
// @Success 202 {object} responses.UserResponse "Retorna um usuário atualizado"
// @Failure 400 {object} responses.ErrorResponse "Dados de entrada inválido"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Router /users [put]
func UserUpdateController(c *gin.Context) {
	services.UserUpdateService(c)
}

// UpdatePasswordUser godoc
// @Summary Atualizar a senha do usuário
// @Description Responsável por atualizar a senha de um usuário
// @Tags Usuários
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param user body requests.UserUpdatePasswordRequest true "Dados do usuário"
// @Success 202 {object} responses.UserResponse "Retorna um usuário atualizado"
// @Failure 400 {object} responses.ErrorResponse "Dados de entrada inválido | Senhas não conferem"
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro interno"
// @Router /users/update-password [put]
func UserUpdatePasswordController(c *gin.Context) {
	services.UserUpdatePasswordService(c)
}

// UploadAvatar godoc
// @Summary Faz upload do avatar do usuário
// @Description Responsável por atualizar o avatar do usuário
// @Tags Usuários
// @Security ApiKeyAuth
// @Accept mpfd
// @Produce json
// @Param avatar formData file true "Arquivo de imagem do usuário (JPEG, PNG, GIF, máx 5MB)"
// @Success 202 {object} responses.UserResponse "Retorna um usuário com o avatar atualizado"
// @Failure 400 {object} responses.ErrorResponse "Nenhum arquivo 'avatar' fornecido | Formato de arquivo inválido | "
// @Failure 401 {object} responses.ErrorResponse "Não autenticado"
// @Failure 429 "Limite de requisições excedido"
// @Failure 500 {object} responses.ErrorResponse "Erro ao tentar salvar o arquivo"
// @Router /users [patch]
func UserUploadController(c *gin.Context) {
	services.UserUploadService(c)
}

func UserIndexController(c *gin.Context) {
	services.UserIndexService(c)
}

func UserDestroyController(c *gin.Context) {
	services.UserDestroyService(c)
}

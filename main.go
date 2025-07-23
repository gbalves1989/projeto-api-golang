package main

import (
	"os"

	"github.com/gbalves1989/projeto-api-golang/config"
	"github.com/gbalves1989/projeto-api-golang/database"
	"github.com/gbalves1989/projeto-api-golang/routes"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"

	_ "github.com/gbalves1989/projeto-api-golang/docs"
)

//	@title			API de estoque de produtos
//	@version		1.0
//	@description	API desenvolvida em golang para cadastrar categoria e produtos
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host						localhost:8080
// @BasePath					/api/v1
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	config.LoadConfig()
	utils.Log(utils.INFO, "Aplicação iniciada...")

	database.ConnectDatabase(false)
	if err := database.MigrateModels(); err != nil {
		utils.Log(utils.ERROR, "Falha na auto-migração do banco de dados: %v", err)
		panic("Falha na auto-migração do banco de dados")
	}
	utils.Log(utils.INFO, "Migrações do banco de dados executadas.")

	if _, err := os.Stat(config.AppConfig.UploadDir); os.IsNotExist(err) {
		os.MkdirAll(config.AppConfig.UploadDir, os.ModePerm)
		utils.Log(utils.INFO, "Diretório de uploads criado: %s", config.AppConfig.UploadDir)
	}

	router := gin.Default()
	router.Static("/uploads", config.AppConfig.UploadDir)
	routes.SetupRoutes(router, database.DB)

	port := config.AppConfig.Port
	router.Run(":" + port)
}

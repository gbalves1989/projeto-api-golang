package database

import (
	"log"
	"sync"

	"github.com/gbalves1989/projeto-api-golang/config"
	"github.com/gbalves1989/projeto-api-golang/database/models"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func ConnectDatabase(testMode bool) {
	once.Do(func() {
		var dsn string

		if testMode {
			dsn = config.AppConfig.DatabaseTestURL
		} else {
			dsn = config.AppConfig.DatabaseURL
		}

		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			utils.Log(utils.ERROR, "Ops! Aplicação encerrada...")
			log.Fatalf("Falha ao conectar ao banco de dados (%s): %v", dsn, err)
		}

		utils.Log(utils.INFO, "Conexão com o banco de dados (%s) estabelecida com sucesso!\n", dsn)
	})
}

func MigrateModels() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
	)
}

func TruncateTables() {
	if err := DB.Exec("TRUNCATE TABLE products RESTART IDENTITY CASCADE").Error; err != nil {
		utils.Log(utils.ERROR, "Erro ao tentar remover todas as linhas da tabela de produtos.")
	}

	if err := DB.Exec("TRUNCATE TABLE categories RESTART IDENTITY CASCADE").Error; err != nil {
		utils.Log(utils.ERROR, "Erro ao tentar remover todas as linhas da tabela de categoria.")
	}

	if err := DB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error; err != nil {
		utils.Log(utils.ERROR, "Erro ao tentar remover todas as linhas da tabela de usu.")
	}
}

package config

import (
	"log"
	"os"
	"strings"

	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL     string
	DatabaseTestURL string
	JWTSecret       string
	UploadDir       string
	TestUploadDir   string
	Port            string
	RateLimit       string
	CORSOrigins     []string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		utils.Log(utils.ERROR, "Ops! Aplicação encerrada...")
		log.Fatal("ERROR: Arquivo de configuração não foi encontrado.")
	}

	AppConfig.DatabaseURL = getEnv("DATABASE_URL")
	AppConfig.DatabaseTestURL = getEnv("DATABASE_TEST_URL")
	AppConfig.UploadDir = getEnvDefault("UPLOAD_DIR", "./uploads")
	AppConfig.JWTSecret = getEnv("JWT_SECRET")
	AppConfig.TestUploadDir = getEnvDefault("TEST_UPLOAD_DIR", "./test_uploads")
	AppConfig.Port = getEnvDefault("PORT", "8080")
	AppConfig.RateLimit = getEnvDefault("RATE_LIMIT", "100-H")
	AppConfig.CORSOrigins = parseCORSOrigins(getEnvDefault("CORS_ORIGINS", "*"))

	if AppConfig.DatabaseURL == "" {
		utils.Log(utils.ERROR, "Ops! Aplicação encerrada...")
		log.Fatal("ERROR: DATABASE_URL não definida nas variáveis de ambiente.")
	}

	if AppConfig.DatabaseTestURL == "" {
		utils.Log(utils.ERROR, "Ops! Aplicação encerrada...")
		log.Fatal("Erro: DATABASE_TEST_URL não definida nas variáveis de ambiente.")
	}

	if AppConfig.JWTSecret == "" {
		utils.Log(utils.ERROR, "Ops! Aplicação encerrada...")
		log.Fatal("Erro: JWT_SECRET não definida nas variáveis de ambiente.")
	}

	utils.Log(utils.INFO, "Configurações carregadas com sucesso!")
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		utils.Log(utils.ERROR, "Ops! Aplicação encerrada...")
		log.Fatalf("ERROR: Variável de ambiente %s não definida.", key)
	}

	return value
}

func getEnvDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func parseCORSOrigins(origins string) []string {
	if origins == "*" {
		return []string{"*"}
	}

	splitOrigins := strings.Split(origins, ",")
	trimmedOrigins := make([]string, 0)
	for _, o := range splitOrigins {
		trimmed := strings.TrimSpace(o)
		if trimmed != "" {
			trimmedOrigins = append(trimmedOrigins, trimmed)
		}
	}

	return trimmedOrigins
}

package middlewares

import (
	"github.com/gbalves1989/projeto-api-golang/config"
	"github.com/gbalves1989/projeto-api-golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func RateLimitMiddleware() gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted(config.AppConfig.RateLimit) // Ex: "100-H", "10-M"

	if err != nil {
		utils.Log(utils.ERROR, "Erro ao parsear a taxa de limite.")
		panic("Erro ao configurar o rate limiter: " + err.Error())
	}

	store := memory.NewStore()
	middleware := mgin.NewMiddleware(limiter.New(store, rate))

	return func(c *gin.Context) {
		middleware(c)
	}
}

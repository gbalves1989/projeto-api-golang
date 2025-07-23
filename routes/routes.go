package routes

import (
	"github.com/gbalves1989/projeto-api-golang/http/controllers"
	"github.com/gbalves1989/projeto-api-golang/http/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authorized := router.Group("/api")
	{
		authGroup := authorized.Group("/")
		authGroup.Use(middlewares.RateLimitMiddleware())
		authGroup.POST("/register", controllers.AuthRegisterController)
		authGroup.POST("/login", controllers.AuthLoginController)
		authGroup.DELETE("/delete", controllers.UserDestroyController)

		userGroup := authorized.Group("/v1/users")
		userGroup.Use(middlewares.AuthMiddleware())
		userGroup.Use(middlewares.RateLimitMiddleware())
		userGroup.GET("/me", controllers.UserMeController)
		userGroup.GET("/", controllers.UserIndexController)
		userGroup.PUT("/", controllers.UserUpdateController)
		userGroup.PUT("/update-password", controllers.UserUpdatePasswordController)
		userGroup.PATCH("/", controllers.UserUploadController)

		categoryGroup := authorized.Group("/v1/categories")
		categoryGroup.Use(middlewares.AuthMiddleware())
		categoryGroup.Use(middlewares.RateLimitMiddleware())
		categoryGroup.POST("/", controllers.CategoryStoreController)
		categoryGroup.GET("/:id", controllers.CategoryShowController)
		categoryGroup.GET("/", controllers.CategoryIndexController)
		categoryGroup.PUT("/:id", controllers.CategoryUpdateController)
		categoryGroup.DELETE("/:id", controllers.CategoryDestroyController)

		productGroup := authorized.Group("/v1/products")
		productGroup.Use(middlewares.AuthMiddleware())
		productGroup.Use(middlewares.RateLimitMiddleware())
		productGroup.POST("/", controllers.ProductStoreController)
		productGroup.GET("/", controllers.ProductIndexController)
		productGroup.GET("/:id", controllers.ProductShowController)
		productGroup.PUT("/:id", controllers.ProductUpdateController)
		productGroup.PATCH("/upload/:id", controllers.ProductUploadController)
		productGroup.DELETE("/:id", controllers.ProductDestroyController)
	}
}

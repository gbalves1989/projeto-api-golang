package repositories

import (
	"github.com/gbalves1989/projeto-api-golang/database"
	"github.com/gbalves1989/projeto-api-golang/database/models"
	"github.com/gbalves1989/projeto-api-golang/http/requests"
)

func ProductShowByNameRepository(name string) models.Product {
	var product models.Product
	database.DB.Where("name = ?", name).First(&product)
	return product
}

func ProductStoreRepository(input requests.ProductCreateOrUpdateRequest) models.Product {
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
		CategoryID:  input.CategoryID,
	}

	database.DB.Create(&product)
	database.DB.Preload("Category").First(&product, product.ID)
	return product
}

func CategoryByProductRepository(product models.Product) models.Product {
	database.DB.Preload("Category").First(&product, product.ID)
	return product
}

func CountProductsRepository() int64 {
	var total int64
	database.DB.Model(&models.Product{}).Count(&total)
	return total
}

func ProductIndexRepository(pagination requests.PaginationRequest) []models.Product {
	var products []models.Product
	database.DB.Limit(pagination.Limit).Offset(pagination.Offset).Find(&products)
	return products
}

func ProductShowRepository(productID uint64) models.Product {
	var product models.Product
	database.DB.First(&product, productID)
	return product
}

func ProductUpdateRepository(input requests.ProductCreateOrUpdateRequest, product models.Product) models.Product {
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	product.CategoryID = input.CategoryID

	database.DB.Save(&product)
	return product
}

func ProductUploadRepository(filePath string, product models.Product) models.Product {
	product.ImageURL = filePath
	database.DB.Save(&product)
	return product
}

func ProductDestroyRepository(product models.Product) {
	database.DB.Delete(&product)
}

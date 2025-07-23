package repositories

import (
	"github.com/gbalves1989/projeto-api-golang/database"
	"github.com/gbalves1989/projeto-api-golang/database/models"
	"github.com/gbalves1989/projeto-api-golang/http/requests"
)

func CategoryShowByNameRepository(name string) models.Category {
	var category models.Category
	database.DB.Where("name = ?", name).First(&category)
	return category
}

func CategoryStoreRepository(input requests.CategoryCreateOrUpdateRequest) models.Category {
	category := models.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	database.DB.Create(&category)
	return category
}

func CategoryShowRepository(categoryID uint64) models.Category {
	var category models.Category
	database.DB.First(&category, categoryID)
	return category
}

func CountCategoriesRepository() int64 {
	var total int64
	database.DB.Model(&models.Category{}).Count(&total)
	return total
}

func CategoryIndexRepository(pagination requests.PaginationRequest) []models.Category {
	var categories []models.Category
	database.DB.Limit(pagination.Limit).Offset(pagination.Offset).Find(&categories)
	return categories
}

func CategoryUpdateRepository(input requests.CategoryCreateOrUpdateRequest, category models.Category) models.Category {
	category.Name = input.Name
	category.Description = input.Description
	database.DB.Save(&category)
	return category
}

func CategoryDestroyRepository(category models.Category) {
	database.DB.Delete(&category)
}

func ProductsByCategoryRepository(category models.Category, categoryID uint64) models.Category {
	database.DB.Preload("Products").First(&category, categoryID)
	return category
}

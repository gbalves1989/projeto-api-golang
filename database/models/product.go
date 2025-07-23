package models

import (
	"github.com/gbalves1989/projeto-api-golang/http/responses"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description,omitempty"`
	Price       float64  `json:"price" binding:"required,min=0"`
	Stock       int      `json:"stock" binding:"required,min=0"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category"`
	ImageURL    string   `json:"image_url,omitempty"`
}

func (p *Product) ToProductResponse() responses.ProductData {
	return responses.ProductData{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		CategoryID:  p.CategoryID,
		ImageURL:    p.ImageURL,
	}
}

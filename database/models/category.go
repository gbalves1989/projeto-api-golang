package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string    `gorm:"uniqueIndex" json:"name" binding:"required"`
	Description string    `json:"description,omitempty"`
	Products    []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}

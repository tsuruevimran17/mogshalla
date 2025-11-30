package models

import "gorm.io/gorm"

type Product struct {
	*gorm.Model `json:"-"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	InStock     bool      `json:"in_stock"`
	CategoryID  uint      `json:"category_id"`
	Category    *Category `json:"category"`
}

type CreateProduct struct {
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	InStock     bool      `json:"in_stock"`
	CategoryID  uint      `json:"category_id"`
	Category    *Category `json:"category"`
}

type UpdateProduct struct {
	Name        *string   `json:"name"`
	Price       *int      `json:"price"`
	Description *string   `json:"description"`
	InStock     *bool     `json:"in_stock"`
	CategoryID  *uint     `json:"category_id"`
	Category    *Category `json:"category"`
}

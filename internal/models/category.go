package models

import "gorm.io/gorm"

type Category struct {
	*gorm.Model `json:"-"`
	Name        string    `json:"name"`
	Products    []Product `json:"-"`
}

type CreateCategory struct {
	Name string `json:"name"`
}

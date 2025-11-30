package repository

import (
	"errors"
	"mogshalla/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	Get(categories *[]models.Category) error
	Delete(id uint) error
}

type gormCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &gormCategoryRepository{db: db}
}

func (r *gormCategoryRepository) Create(category *models.Category) error {

	if category == nil {
		return errors.New("category is nil")
	}

	return r.db.Save(category).Error
}

func (r *gormCategoryRepository) Get(categories *[]models.Category)  error {
	

	if err := r.db.Find(&categories).Error; err != nil {

		return  errors.New("категории не найдены")

	}

	return nil
}

func (r *gormCategoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}

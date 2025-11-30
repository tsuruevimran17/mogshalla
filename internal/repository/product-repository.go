package repository

import (
	"mogshalla/internal/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	Update(product *models.Product) error
	Get(categoryID uint) ([]models.Product, error)
	GetById(id uint) (*models.Product, error)
	Delete(id uint) error
}

type gormProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &gormProductRepository{db: db}
}

func (r *gormProductRepository) Create(product *models.Product) error {

	if product == nil {
		return nil
	}

	return r.db.Preload("Category").Create(product).Error
}

func (r *gormProductRepository) Update(product *models.Product) error {

	if product == nil {
		return nil
	}

	return r.db.Save(product).Error
}

func (r *gormProductRepository) Get(categoryID uint) ([]models.Product, error) {
	var products []models.Product

	if err := r.db.Preload("Category").Where("category_id=?", categoryID).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *gormProductRepository) GetById(id uint) (*models.Product, error) {

	var product models.Product

	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *gormProductRepository) Delete(id uint) error {

	return r.db.Delete(&models.Product{}, id).Error

}

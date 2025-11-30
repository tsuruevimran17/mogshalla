package service

import (
	"errors"
	"mogshalla/internal/models"
	"mogshalla/internal/repository"
)

type ProductService interface {
	Create(req *models.CreateProduct) (*models.Product, error)
	Get(categoryID uint) ([]models.Product, error)
	GetById(id uint) (*models.Product, error)
	Update(id uint, req *models.UpdateProduct) ( error)
	Delete(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(req *models.CreateProduct) (*models.Product, error) {
	if err := s.validate(req); err != nil {
		return nil, errors.New(err.Error())
	}
	product := models.Product{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		InStock:     req.InStock,
		CategoryID:  req.CategoryID,
	}
	if err := s.repo.Create(&product); err != nil {
		return nil, errors.New("ошибка при создании продукта")
	}
	return &product, nil
}

func (s *productService) validate(req *models.CreateProduct) error {
	if req.Name == "" {
		return errors.New("имя не может быть пустым")
	}
	if req.Price <= 0 {
		return errors.New("цена должна быть больше нуля")
	}
	if req.CategoryID == 0 {
		return errors.New("категория не выбрана")
	}
	if req.Description == "" {
		return errors.New("описание не может быть пустым")
	}
	return nil
}

func (s *productService) validateUpdate(req *models.UpdateProduct) error {

	if req.Name != nil && *req.Name == "" {
		return errors.New("имя не может быть пустым")
	}

	if req.Price != nil && *req.Price <= 0 {
		return errors.New("цена должна быть больше нуля")
	}

	if req.CategoryID != nil && *req.CategoryID == 0 {
		return errors.New("категория не выбрана")
	}

	if req.Description != nil && *req.Description == "" {
		return errors.New("описание не может быть пустым")
	}

	return nil
}

func (r *productService) Get(categoryID uint) ([]models.Product, error) {
	products, err := r.repo.Get(categoryID)
	if err != nil {
		return nil, errors.New("ошибка при получении продуктов")
	}
	return products, nil
}
func (r *productService) Delete(id uint) error {
	return r.repo.Delete(id)
}
func (r *productService) GetById(id uint) (*models.Product, error) {
	product, err := r.repo.GetById(id)
	if err != nil {
		return nil, errors.New("продукт не найден")
	}
	return product, nil
}
func (r *productService) Update(id uint, req *models.UpdateProduct) ( error) {
	product, err := r.repo.GetById(id)

	if err != nil {
		return  errors.New("продукт не найден")
	}

	if err := r.validateUpdate(req); err != nil {
		return errors.New("невалидные данные продукта")
	}

	if req.Name != nil {
		product.Name = *req.Name
	}

	if req.Price != nil {
		product.Price = *req.Price
	}

	if req.Description != nil {
		product.Description = *req.Description
	}

	if req.InStock != nil {
		product.InStock = *req.InStock
	}

	if req.CategoryID != nil {
		product.CategoryID = *req.CategoryID
	}

	if err := r.repo.Update(product); err != nil {
		return errors.New("ошибка при обновлении продукта")
	}

	return  nil
}

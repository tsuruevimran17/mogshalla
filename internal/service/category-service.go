package service

import (
	"errors"
	"mogshalla/internal/models"
	"mogshalla/internal/repository"
)

type CategoryService interface {
	Create(req models.CreateCategory)(*models.Category, error)
	Get() ([]models.Category, error)
	Delete(id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) Create(req models.CreateCategory) (*models.Category, error) {

	if err := s.validate(req); err != nil {
		return nil, errors.New("невалидные данные категории")
	}

	category := models.Category{
		Name: req.Name,
	}

	if err := s.repo.Create(&category); err != nil {

		return nil, errors.New("ошибка при создании категории")

	}

	return &category, nil

}

func (s *categoryService) validate(req models.CreateCategory) error {

	if req.Name == "" {
		return errors.New("имя не может быть пустым")
	}

	return nil
}

func (s *categoryService) Get() ([]models.Category, error) {
	var categories []models.Category

	if err := s.repo.Get(&categories); err != nil {
		return nil, errors.New("ошибка при получении категорий")
	}


	return categories, nil
}

func (s *categoryService) Delete(id uint) error {

	if err := s.repo.Delete(id); err != nil {
		return errors.New("ошибка при удалении категории")
	}

	return nil
}
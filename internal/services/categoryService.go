package services

import (
	"news-portal/internal/models/category"
	model "news-portal/internal/models/category"
	"news-portal/internal/repositories"
	"news-portal/internal/utils"
	"time"

	"github.com/gosimple/slug"
)

// ICategoryService defines the interface for category service.
type ICategoryService interface {
	Slugify(name string) string
	GetAllCategories() ([]model.Category, error)
	GetCategoryBySlug(slug string) (*model.Category, error)
	GetCategoryByID(id uint) (*model.Category, error)
	CreateCategory(data category.Category) (model.Category, error)
	UpdateCategory(id uint, data category.Category) (*model.Category, error)
	DeleteCategory(id uint) error
}

// CategoryService provides operations on categories.
type CategoryService struct {
	repository repositories.CategoryRepository
}

// NewCategoryService creates a new CategoryService.
func NewCategoryService(repository repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repository: repository}
}

// Slugify generates a slug from the given name.
func (s *CategoryService) Slugify(name string) string {
	return slug.Make(name)
}

// CreateCategory creates a new category.
func (s *CategoryService) CreateCategory(data category.Category) (model.Category, error) {
	data.Slug = utils.Slugify(data.Name)
	data.CreatedAt = time.Now()
	err := s.repository.Create(&data)
	if err != nil {
		return category.Category{}, err
	}
	return data, nil
}

// GetAllCategories returns all categories.
func (s *CategoryService) GetAllCategories() ([]model.Category, error) {
	return s.repository.FindAll()
}

func (s *CategoryService) GetCategoryBySlug(slug string) (*model.Category, error) {
	return s.repository.FindBySlug(slug)
}

func (s *CategoryService) GetCategoryByID(id uint) (*model.Category, error) {
	return s.repository.FindByID(id)
}

func (s *CategoryService) UpdateCategory(id uint, data category.Category) (*model.Category, error) {
	category, err := s.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	category.Name = data.Name
	category.Slug = utils.Slugify(data.Name)
	category.UpdatedAt = time.Now()
	err = s.repository.Update(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.repository.Delete(id)
}

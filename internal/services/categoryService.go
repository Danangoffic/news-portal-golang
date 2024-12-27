package services

import (
	model "news-portal/internal/models/category"
	"news-portal/internal/repositories"
	"news-portal/internal/utils"
	"time"

	"github.com/gosimple/slug"
)

// ICategoryService defines the interface for category service.
type ICategoryService interface {
	Slugify(name string) string
	CreateCategory(name string) *model.Category
	GetAllCategories() ([]model.Category, error)
	GetCategoryBySlug(slug string) (*model.Category, error)
	GetCategoryByID(id uint) (*model.Category, error)
	UpdateCategory(id uint, name string) (*model.Category, error)
	DeleteCategory(id uint) error
}

// CategoryService provides operations on categories.
type CategoryService struct {
	repository repositories.ICategoryRepository
}

// NewCategoryService creates a new CategoryService.
func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

// Slugify generates a slug from the given name.
func (s *CategoryService) Slugify(name string) string {
	return slug.Make(name)
}

// CreateCategory creates a new category.
func (s *CategoryService) CreateCategory(name string) *model.Category {
	slug := utils.Slugify(name)
	category := &model.Category{
		Name:      name,
		Slug:      slug,
		CreatedAt: time.Now(),
	}
	err := s.repository.Create(category)
	if err != nil {
		return nil
	}
	return category
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

func (s *CategoryService) UpdateCategory(id uint, name string) (*model.Category, error) {
	category, err := s.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	category.Name = name
	category.Slug = utils.Slugify(name)
	err = s.repository.Update(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.repository.Delete(id)
}

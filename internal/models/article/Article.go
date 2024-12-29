package article

import (
	"news-portal/internal/models/category"
	"time"
)

type ArticleStatus string

const (
	Draft     ArticleStatus = "draft"
	Published ArticleStatus = "published"
	Archived  ArticleStatus = "archived"
)

type Article struct {
	ID         int               `json:"id" gorm:"primary_key"`
	Title      string            `json:"title"`
	Slug       string            `json:"slug"`
	Content    string            `json:"content"`
	Author     string            `json:"author"`
	Status     ArticleStatus     `json:"status"`
	CategoryID uint              `json:"category_id"`
	Category   category.Category `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

type CreateArticle struct {
	Title      string `json:"title" validate:"required,max=255"`
	Content    string `json:"content" validate:"required"`
	Author     string `json:"author" validate:"required,max=255"`
	Status     string `json:"status" validate:"required"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

type UpdateArticle struct {
	Title      string `json:"title" validate:"required,max=255"`
	Content    string `json:"content" validate:"required"`
	Author     string `json:"author" validate:"required,max=255"`
	Status     string `json:"status" validate:"required"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

type UpdateArticleStatus struct {
	Status string `json:"status" validate:"required"`
}

type ArticleResponse struct {
	Data   []Article `json:"data"`
	Status string    `json:"status"`
}

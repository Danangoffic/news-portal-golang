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
	CategoryID int               `json:"category_id"`
	Category   category.Category `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

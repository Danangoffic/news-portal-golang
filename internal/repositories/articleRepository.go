package repositories

import (
	model "news-portal/internal/models/article"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	GetArticles() ([]model.Article, error)
	GetArticleByID(id uint) (*model.Article, error)
	GetArticlesByStatus(status string) ([]model.Article, error)
	GetArticleBySlug(slug string) (*model.Article, error)
	CreateArticle(article *model.Article) error
	UpdateArticle(article model.Article) error
	DeleteArticle(article *model.Article) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

func (r *articleRepository) GetArticles() ([]model.Article, error) {
	var articles []model.Article
	if err := r.db.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) GetArticleByID(id uint) (*model.Article, error) {
	var article model.Article
	if err := r.db.First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *articleRepository) GetArticlesByStatus(status string) ([]model.Article, error) {
	var articles []model.Article
	if err := r.db.Where("status = ?", status).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) GetArticleBySlug(slug string) (*model.Article, error) {
	var article model.Article
	if err := r.db.Where("slug = ?", slug).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *articleRepository) CreateArticle(article *model.Article) error {
	if err := r.db.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) UpdateArticle(article model.Article) error {
	if err := r.db.Save(article).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) DeleteArticle(article *model.Article) error {
	if err := r.db.Delete(article).Error; err != nil {
		return err
	}
	return nil
}

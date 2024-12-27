package services

import (
	model "news-portal/internal/models/article"
	"news-portal/internal/repositories"
	"news-portal/internal/utils"
)

type ArticleService interface {
	GetArticleById(id uint) (*model.Article, error)
	GetArticles() ([]model.Article, error)
	GetArticlesByStatus(status string) ([]model.Article, error)
	GetArticleBySlug(slug string) (*model.Article, error)
	CreateArticle(article *model.Article) error
	UpdateArticle(article *model.Article) error
	DeleteArticle(article *model.Article) error
}

type articleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

func (s *articleService) GetArticleById(id uint) (*model.Article, error) {
	return s.repo.GetArticleByID(id)
}
func (s *articleService) GetArticles() ([]model.Article, error) {
	return s.repo.GetArticles()
}

func (s *articleService) GetArticlesByStatus(status string) ([]model.Article, error) {
	return s.repo.GetArticlesByStatus(status)
}

func (s *articleService) GetArticleBySlug(slug string) (*model.Article, error) {
	return s.repo.GetArticleBySlug(slug)
}

func (s *articleService) CreateArticle(article *model.Article) error {
	article.Slug = utils.Slugify(article.Title)
	article.Status = model.Draft

	return s.repo.CreateArticle(article)
}

func (s *articleService) UpdateArticle(article *model.Article) error {
	return s.repo.UpdateArticle(article)
}

func (s *articleService) DeleteArticle(article *model.Article) error {
	return s.repo.DeleteArticle(article)
}

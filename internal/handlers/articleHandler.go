package handlers

import (
	"net/http"
	"news-portal/internal/models/article"
	"news-portal/internal/services"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

var lock sync.Mutex

type ArticleHandler struct {
	service services.ArticleService
}

func NewArticleHandler(service services.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: service}
}

func (h *ArticleHandler) GetAllArticles(c echo.Context) error {
	articles, err := h.service.GetArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error(), "status": "failed"})
	}

	return c.JSON(http.StatusOK, map[string]any{"data": articles, "status": "success"})
}

func (h *ArticleHandler) GetArticleById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	article, err := h.service.GetArticleById(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error(), "status": "failed"})
	}

	return c.JSON(http.StatusOK, map[string]any{"data": article, "status": "success"})
}

func (h *ArticleHandler) GetArticlesByStatus(c echo.Context) error {
	status := c.QueryParam("status")
	if status == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Status query parameter is required"})
	}

	articles, err := h.service.GetArticlesByStatus(status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error(), "status": "failed"})
	}

	return c.JSON(http.StatusOK, map[string]any{"data": articles, "status": "success"})
}

func (h *ArticleHandler) GetArticleBySlug(c echo.Context) error {
	slug := c.Param("slug")
	if slug == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Slug parameter is required"})
	}

	article, err := h.service.GetArticleBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error(), "status": "failed"})
	}

	return c.JSON(http.StatusOK, map[string]any{"data": article, "status": "success"})
}

func (h *ArticleHandler) CreateArticle(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	var createArticleData article.CreateArticle
	if err := c.Bind(&createArticleData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid request payload", "status": "failed"})
	}
	article := article.Article{
		Title:      createArticleData.Title,
		Content:    createArticleData.Content,
		Author:     createArticleData.Author,
		Status:     article.ArticleStatus(createArticleData.Status),
		CategoryID: createArticleData.CategoryID,
	}
	err := h.service.CreateArticle(&article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error(), "status": "failed"})
	}

	return c.JSON(http.StatusCreated, map[string]any{"data": article, "status": "success"})
}

func (h *ArticleHandler) UpdateArticle(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var article article.Article
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid request payload", "status": "failed"})
	}

	ID := int(id)
	err = h.service.UpdateArticle(ID, article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error(), "status": "failed"})
	}

	return c.JSON(http.StatusOK, map[string]any{"data": article, "status": "success"})
}

func (h *ArticleHandler) DeleteArticle(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	article := &article.Article{ID: int(id)}
	err = h.service.DeleteArticle(article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error(), "status": "failed"})
	}

	return c.JSON(http.StatusOK, map[string]any{"status": "success"})
}

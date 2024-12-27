package handlers

import (
	"log"
	"net/http"
	"news-portal/internal/models/category"
	"news-portal/internal/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	Service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: service}
}

func (h *CategoryHandler) GetCategories(c echo.Context) error {
	categories, err := h.Service.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]any{
		"data": categories, "status": "success",
	})
}

func (h *CategoryHandler) GetCategory(c echo.Context) error {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid ID"})

	}
	category, err := h.Service.GetCategoryByID(uint(uintID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]any{
		"data": category, "status": "success",
	})
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	var category category.Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": err.Error()})

	}
	if _, err := h.Service.CreateCategory(category); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})

	}
	return c.NoContent(http.StatusCreated)
}

func (h *CategoryHandler) UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	log.Println("update category with id : ", id)
	var category category.Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": err.Error()})

	}
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid ID"})

	}
	if _, err := h.Service.UpdateCategory(uint(uintID), category); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})

	}
	return c.NoContent(http.StatusOK)
}

func (h *CategoryHandler) DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid ID"})

	}
	if err := h.Service.DeleteCategory(uint(uintID)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})

	}
	return c.NoContent(http.StatusOK)
}

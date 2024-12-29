package handlers

import (
	"net/http"
	"news-portal/internal/models/user"
	"news-portal/internal/services"
	"news-portal/pkg/auth"

	"github.com/labstack/echo/v4"
)

type handler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *handler {
	return &handler{userService}
}

func (h *handler) ProtectedHandler(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header is missing"})
	}

	tokenString := authHeader[len("Bearer "):] // Hapus "Bearer " dari header

	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}

	// Sekarang Anda memiliki claims, Anda bisa mengakses informasi pengguna
	// Contoh:
	userID := claims.UserID
	username := claims.Username

	// Lakukan sesuatu dengan informasi pengguna
	_ = userID
	_ = username
	c.Set("userID", userID)
	c.Set("username", username)

	return c.JSON(http.StatusOK, map[string]string{"message": "Protected resource accessed"})
}

func (h *handler) LoginHandler(c echo.Context) error {
	var loginRequest user.LoginUserRequest
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	var userData *user.User
	userData, err := h.userService.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user"})
	}

	token, err := auth.GenerateToken(userData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (h *handler) RefreshTokenHandler(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header is missing"})
	}

	tokenString := authHeader[len("Bearer "):]

	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}

	// Generate a new token with the same claims
	newToken, err := auth.GenerateToken(&user.User{
		ID:       claims.UserID,
		Username: claims.Username,
		Email:    claims.Email,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate new token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": newToken})
}

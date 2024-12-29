package AuthJWTNews

import (
	"net/http"
	"news-portal/pkg/auth"

	"github.com/labstack/echo/v4"
)

func AuthJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header is missing"})
		}

		tokenString := authHeader[len("Bearer "):]

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
		}

		// Simpan userID dan username ke dalam context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		return next(c) // Panggil handler selanjutnya dalam chain
	}
}

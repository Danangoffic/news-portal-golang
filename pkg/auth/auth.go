package auth

import (
	"fmt"
	"os"
	"time"

	"news-portal/internal/models/jwtPack"
	"news-portal/internal/models/user"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user *user.User) (string, error) {
	// Ambil secret key dari environment variable
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY environment variable not set")
	}
	claims := &jwtPack.Claims{
		Username: user.Username,
		UserID:   user.ID,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token berlaku 24 jam
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwtPack.Claims, error) {
	// Ambil secret key dari environment variable
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY environment variable not set")
	}
	claims := &jwtPack.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Validasi apakah token sudah expired
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token has expired")
	}

	return claims, nil
}

package jwtPack

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	UserID   uint   `json:"user_id"` // Contoh menambahkan ID pengguna
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

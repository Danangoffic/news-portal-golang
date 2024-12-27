package user

import (
	"time"
)

type UserStatus string
type UserType string

const (
	Active    UserStatus = "active"
	Inactive  UserStatus = "inactive"
	Suspended UserStatus = "suspended"
)

const (
	Admin  UserType = "admin"
	Editor UserType = "editor"
	Viewer UserType = "viewer"
)

type User struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	FullName  string     `json:"full_name"`
	Password  string     `json:"password"`
	Status    UserStatus `json:"status"`
	Type      UserType   `json:"type"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

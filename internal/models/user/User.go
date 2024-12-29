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
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	FullName  string     `json:"full_name"`
	Password  string     `json:"password"`
	Status    UserStatus `json:"status"`
	Type      UserType   `json:"type"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Type     string `json:"type" validate:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"full_name" validate:"required"`
	Type     string `json:"type" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type dataUserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Type     string `json:"type"`
	Status   string `json:"status"`
}

type UserResponse struct {
	Data   dataUserResponse `json:"data"`
	Status string           `json:"status"`
}

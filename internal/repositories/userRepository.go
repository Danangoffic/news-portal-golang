package repositories

import (
	"news-portal/internal/models/user"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *user.User) error
	GetUserByID(id uint) (*user.User, error)
	GetUserByEmail(email string) (*user.User, error)
	UpdateUser(user *user.User) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id uint) (*user.User, error) {
	var user user.User
	err := r.db.First(&user, id).Error // Menggunakan First untuk mencari berdasarkan ID
	return &user, err
}

func (r *userRepository) GetUserByEmail(email string) (*user.User, error) {
	var user user.User
	err := r.db.Where("email = ?", email).First(&user).Error // Menggunakan Where dan First untuk mencari berdasarkan email
	return &user, err
}

func (r *userRepository) UpdateUser(user *user.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&user.User{}, id).Error // Menghapus user berdasarkan ID
}

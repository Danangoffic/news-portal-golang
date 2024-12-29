package category

import (
	"time"
)

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	Slug      string    `json:"slug" gorm:"size:255;not null;unique"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type CreateCategory struct {
	Name string `json:"name" gorm:"size:255"`
}

type UpdateCategory struct {
	Name string `json:"name" gorm:"size:255"`
}

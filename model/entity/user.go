package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	Age       int            `json:"age"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

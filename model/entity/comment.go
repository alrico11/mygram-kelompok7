package entity

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int            `json:"id"`
	Message   string         `json:"message"`
	PhotoID   int            `json:"photo_id"`
	UserID    int            `json:"user_id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

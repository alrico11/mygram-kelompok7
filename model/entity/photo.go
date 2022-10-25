package entity

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID        int            `json:"id"`
	Title     string         `json:"title"`
	Caption   string         `json:"caption"`
	PhotoURL  string         `json:"photo_url"`
	UserID    int            `json:"user_id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

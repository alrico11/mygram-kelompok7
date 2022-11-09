package entity

import (
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	PhotoID   int       `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User
	Photo     Photo
}

func (Comment) TableName() string {
	return "comment"
}

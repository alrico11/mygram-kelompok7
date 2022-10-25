package response

import (
	"project2/model/entity"
	"time"
)

type CreateCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"-"`
}

type GetCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	CreatedAt time.Time `json:"-"`
	Photo     entity.Photo
}

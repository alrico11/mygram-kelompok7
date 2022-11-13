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
	CreatedAt time.Time `json:"created_at"`
}

type GetCommentResponse struct {
	ID        int          `json:"id"`
	Message   string       `json:"message"`
	PhotoID   int          `json:"photo_id"`
	UserID    int          `json:"user_id"`
	UpdatedAt time.Time    `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
	User      CommentUser  `json:"User"`
	Photo     CommentPhoto `json:"Photo"`
}

type CommentUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommentPhoto struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   int    `json:"user_id"`
}

func GetAllComment(comment entity.Comment, photo entity.Photo) GetCommentResponse {
	response := GetCommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
		CreatedAt: comment.CreatedAt,
		User: CommentUser{
			ID:       comment.User.ID,
			Email:    comment.User.Email,
			Username: comment.User.Username,
		},
		Photo: CommentPhoto{
			ID:       comment.Photo.ID,
			Title:    comment.Photo.Title,
			Caption:  comment.Photo.Caption,
			PhotoURL: comment.Photo.PhotoURL,
			UserID:   comment.Photo.UserID,
		},
	}

	return response
}

type CommentUpdateResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentDeleteResponse struct {
	Message string `json:"message"`
}

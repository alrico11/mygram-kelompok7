package response

import (
	"project2/model/entity"
	"time"
)

type CreatePhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetPhotoUser struct {
	ID        int         `json:"id"`
	Title     string      `json:"title"`
	Caption   string      `json:"caption"`
	PhotoURL  string      `json:"photo_url"`
	UserID    int         `json:"user_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	User      UserInPhoto `json:"User"`
}

// type GetPhotoDetailUser struct {
// 	ID        int              `json:"id"`
// 	Title     string           `json:"title"`
// 	Caption   string           `json:"caption"`
// 	PhotoURL  string           `json:"photo_url"`
// 	CreatedAt time.Time        `json:"created_at"`
// 	User      UserInPhoto      `json:"user"`
// 	Comments  []entity.Comment `json:"comments"`
// }

type UserInPhoto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetAllPhotosUser(photos []entity.Photo) []GetPhotoUser {
	if len(photos) == 0 {
		return []GetPhotoUser{}
	}

	var allPhotoUser []GetPhotoUser

	for _, photo := range photos {
		tmpPhoto := GetPhotoUser{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			CreatedAt: photo.CreatedAt,
		}

		allPhotoUser = append(allPhotoUser, tmpPhoto)
	}

	return allPhotoUser
}

type PhotoUpdateResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoDeleteResponse struct {
	Message string `json:"message"`
}

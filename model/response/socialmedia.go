package response

import (
	"project2/model/entity"
	"time"
)

type SocialMediaCreateResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"social_media_url"`
	UsedID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type SocialMediaUpdateResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"social_media_url"`
	UsedID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetSocialMedia struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	URL       string      `json:"social_media_url"`
	UsedID    int         `json:"user_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdateAt  time.Time   `json:"updated_at"`
	User      entity.User `json:"User"`
}

type SocialMediaDeleteResponse struct {
	Message string `json:"message"`
}

func GetAllSocialMedia(social []entity.SocialMedia, user entity.User) []GetSocialMedia {
	if len(social) == 0 {
		return []GetSocialMedia{}
	}

	var allSocialMedia []GetSocialMedia

	for _, socialmedia := range social {
		tmpSocialmedia := GetSocialMedia{
			ID:        socialmedia.ID,
			Name:      socialmedia.Name,
			URL:       socialmedia.URL,
			UsedID:    socialmedia.UserID,
			CreatedAt: socialmedia.CreatedAt,
			User: entity.User{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			},
		}

		allSocialMedia = append(allSocialMedia, tmpSocialmedia)

	}

	return allSocialMedia
}

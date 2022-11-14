package response

import (
	"errors"
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
	UserID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SocialMediaGetResponse struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	URL       string          `json:"social_media_url"`
	UsedID    int             `json:"user_id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdateAt  time.Time       `json:"updated_at"`
	User      SocialMediaUser `json:"User"`
}

type SocialMediaUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type SocialMediaDeleteResponse struct {
	Message string `json:"message"`
}

func GetAllSocialMedia(social []entity.SocialMedia, user entity.User) ([]SocialMediaGetResponse, error) {
	if len(social) == 0 {
		return []SocialMediaGetResponse{}, errors.New("no data")
	}

	var allSocialMedia []SocialMediaGetResponse

	for _, socialmedia := range social {
		tmpSocialmedia := SocialMediaGetResponse{
			ID:        socialmedia.ID,
			Name:      socialmedia.Name,
			URL:       socialmedia.URL,
			UsedID:    socialmedia.UserID,
			CreatedAt: socialmedia.CreatedAt,
			UpdateAt:  socialmedia.UpdatedAt,
			User: SocialMediaUser{
				ID:       socialmedia.User.ID,
				Username: socialmedia.User.Username,
			},
		}

		allSocialMedia = append(allSocialMedia, tmpSocialmedia)

	}

	return allSocialMedia, nil
}

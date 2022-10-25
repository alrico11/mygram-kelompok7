package entity

import "time"

type SosmedCreateesponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"social_media_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type SosmedUpdateResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"social_media_url"`
	UsedID    int       `json:"user_id"`
	UpdatedAt time.Time `json:"date"`
}

type SosmedGetResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"social_media_url"`
	UsedID    int       `json:"user_id"`
	CreatedAt time.Time `json:"date"`
	User      entity.User
}

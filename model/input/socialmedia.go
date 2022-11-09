package input

type SocialInput struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"social_media_url" binding:"required"`
}

type DeleteSocialMedia struct {
	ID int `uri:"id" binding:"required"`
}
type UpdateSocialMedia struct {
	ID int `uri:"id" binding:"required"`
}

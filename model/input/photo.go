package input

type InputPhotos struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required"`
}

type UpdatePhoto struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

type UpdatePhotoIDUser struct {
	ID int `uri:"id" binding:"required"`
}

type DeletePhoto struct {
	ID int `uri:"id" binding:"required"`
}

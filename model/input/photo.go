package input

type PhotoCreateInput struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required"`
}

type PhotoUpdateInput struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}

type PhotoUpdateIDUser struct {
	ID int `uri:"id" binding:"required"`
}

type PhotoDeleteIDUser struct {
	ID int `uri:"id" binding:"required"`
}

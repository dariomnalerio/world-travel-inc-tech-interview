package models

type GetRandomImageResponse struct {
	ImageURL string `json:"image_url"`
	Liked    bool   `json:"liked"`
}

type GetRandomImageRequest struct {
	UserID string `json:"user_id"`
}

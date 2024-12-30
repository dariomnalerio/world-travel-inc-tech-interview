package models

type RequiredUserID struct {
	UserID string `uri:"id" binding:"required,uuid"`
}

type RequiredImageURL struct {
	ImageURL string `uri:"image_url" binding:"required,url"`
}

// Get Image Types
type GetLikedImagesRequest RequiredUserID
type GetLikedImagesResponse struct {
	Images []string `json:"images"`
}

// Like Image types
type LikeImageRequestURL RequiredUserID
type LikeImageRequestBody RequiredImageURL
type LikeImageResponse struct {
	Success bool     `json:"success"`
	Images  []string `json:"images"`
}

// Unlike Image types
type UnlikeImageRequestURL RequiredUserID
type UnlikeImageRequestBody RequiredImageURL
type UnlikeImageResponse struct {
	Success bool `json:"success"`
}

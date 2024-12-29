package repositories

type LikedImagesRepository interface {
	GetLikedImages(userID string) ([]string, error)
	AddLikedImage(userID string, image string) error
	RemoveLikedImage(userID string, imageID string) error
}

type likedImagesRepository struct {
	// map[userID]map[imageURL]struct{} - using struct{} since we only care about existence
	// struct{} is an empty struct that takes up zero bytes of memory, which is useful for sets
	likes map[string]map[string]struct{}
}

func NewLikedImagesRepository() *likedImagesRepository {
	return &likedImagesRepository{
		likes: make(map[string]map[string]struct{}),
	}
}

// AddLikedImage adds an image URL to the list of liked images for a given user.
// If the user does not have any liked images yet, a new entry is created.
//
// Parameters:
//   - userID: The ID of the user liking the image.
//   - imageURL: The URL of the image to be liked.
//
// Returns:
//   - error: An error if the operation fails, otherwise nil.
func (r *likedImagesRepository) AddLikedImage(userID, imageURL string) error {
	if _, exists := r.likes[userID]; !exists {
		r.likes[userID] = make(map[string]struct{})
	}
	r.likes[userID][imageURL] = struct{}{}
	return nil
}

// RemoveLikedImage removes the like for a given image URL by a specific user.
// If the user has not liked the image, the function does nothing.
// Parameters:
//   - userID: The ID of the user unliking the image.
//   - imageURL: The URL of the image to be unliked.
//
// Returns:
//   - error: An error if any issues occur during retrieval
func (r *likedImagesRepository) RemoveLikedImage(userID, imageURL string) error {
	if _, exists := r.likes[userID]; exists {
		delete(r.likes[userID], imageURL)
	}
	return nil
}

// GetLikedImages retrieves a list of image URLs liked by a specific user.
// It returns a slice of image URLs and an error if any occurred.
//
// Parameters:
//   - userID: The ID of the user whose liked images are to be retrieved.
//
// Returns:
//   - []string: A slice of image URLs liked by the user.
//   - error: An error if any issues occur during retrieval.
func (r *likedImagesRepository) GetLikedImages(userID string) ([]string, error) {
	if userLikes, exists := r.likes[userID]; exists {
		images := make([]string, 0, len(userLikes))
		for url := range userLikes {
			images = append(images, url)
		}
		return images, nil
	}
	return []string{}, nil
}

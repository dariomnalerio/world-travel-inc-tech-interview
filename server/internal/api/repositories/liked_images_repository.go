package repositories

import (
	"database/sql"
	"server/db/queries"
	e "server/internal/errors"
)

type LikedImagesRepository interface {
	GetLikedImages(userID string) ([]string, error)
	AddLikedImage(userID string, image string) error
	RemoveLikedImage(userID string, imageID string) error
}

type likedImagesRepository struct {
	db *sql.DB
}

func NewLikedImagesRepository(db *sql.DB) *likedImagesRepository {
	return &likedImagesRepository{
		db: db,
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
	exists, err := queries.GetLikedImage(r.db, userID, imageURL)

	if err != nil {
		return err
	}

	if exists {
		return e.NewError(e.ValidationErr, e.ImageAlreadyLiked, "image already liked", nil)
	}

	err = queries.AddLikedImage(r.db, userID, imageURL)

	if err != nil {
		return e.NewError(e.InternalErr, e.DatabaseError, "failed to add liked image", err)
	}

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
	err := queries.RemoveLikedImage(r.db, userID, imageURL)

	if err != nil {
		return e.NewError(e.InternalErr, e.DatabaseError, "failed to remove liked image", err)
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
	imgs, err := queries.GetLikedImages(r.db, userID)

	if err != nil {
		return nil, e.NewError(e.InternalErr, e.DatabaseError, "failed to get liked images", err)
	}

	return imgs, nil
}

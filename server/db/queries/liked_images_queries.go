package queries

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// AddLikedImage adds an image URL to the list of liked images for a given user.
func AddLikedImage(db *sql.DB, userID, imageURL string) error {
	_, err := db.Exec("INSERT INTO liked_images (user_id, image_url) VALUES ($1, $2)", userID, imageURL)
	if err != nil {
		return err
	}
	return nil
}

// RemoveLikedImage removes the like for a given image URL by a specific user.
func RemoveLikedImage(db *sql.DB, userID, imageURL string) error {
	log.Println("Removing liked image")
	_, err := db.Exec("DELETE FROM liked_images WHERE user_id = $1 AND image_url = $2", userID, imageURL)
	if err != nil {
		return err
	}
	return nil
}

// GetLikedImages retrieves a list of image URLs liked by a specific user.
func GetLikedImages(db *sql.DB, userID string) ([]string, error) {
	var images []string
	rows, err := db.Query("SELECT image_url FROM liked_images WHERE user_id = $1", userID)
	if err == sql.ErrNoRows {
		return []string{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var imageURL string
		if err := rows.Scan(&imageURL); err != nil {
			return nil, err
		}
		images = append(images, imageURL)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return images, nil
}

// GetLikedImage reports whether a user has liked a specific image.
func GetLikedImage(db *sql.DB, userID, imageURL string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM liked_images WHERE user_id = $1 AND image_url = $2)", userID, imageURL).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

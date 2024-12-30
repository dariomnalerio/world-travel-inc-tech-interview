package repositories

import (
	"encoding/json"
	"errors"
	"net/http"
)

// DogRepository defines the interface for dog-related operations
type DogRepository interface {
	GetRandomPicture() (string, error)
}

// DogAPIResponse represents the response from the Dog API
type DogAPIResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// dogAPIRepository is a repository that interacts with an external dog API.
//
// It contains the base URL of the API and an HTTP client to make requests.
type dogAPIRepository struct {
	baseURL string
	client  *http.Client
}

// NewDogAPIRepository creates a new instance of DogRepository with the specified base URL.
//
// It initializes an HTTP client for making API requests.
func NewDogAPIRepository(baseURL string) DogRepository {
	return &dogAPIRepository{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

// GetRandomPicture fetches a random dog picture from the Dog API.
//
// It returns the URL of the picture as a string and an error if any occurred during the process.
func (r *dogAPIRepository) GetRandomPicture() (string, error) {
	randomImagePath := "/breeds/image/random"
	req, err := http.NewRequest(http.MethodGet, r.baseURL+randomImagePath, nil)
	if err != nil {
		return "", err
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var apiResponse DogAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return "", err
	}

	if apiResponse.Status != "success" {
		return "", errors.New("failed to fetch random dog picture")
	}

	return apiResponse.Message, nil

}

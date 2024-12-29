package testing

import (
	e "server/internal/errors"

	"github.com/stretchr/testify/mock"
)

type MockLikedImagesBuilder struct {
	mock        *MockLikedImagesRepository
	likedImages map[string][]string
}

func NewLikedImagesMockBuilder() *MockLikedImagesBuilder {
	return &MockLikedImagesBuilder{
		mock:        &MockLikedImagesRepository{},
		likedImages: make(map[string][]string),
	}
}

func (b *MockLikedImagesBuilder) WithInitialLikedImages(userLikes map[string][]string) *MockLikedImagesBuilder {
	b.likedImages = userLikes
	return b
}

// WithLikedImage sets up the mock to return a successful response when liking an image.
func (b *MockLikedImagesBuilder) WithAddLikedImage(userID, imageURL string) *MockLikedImagesBuilder {
	b.mock.On("AddLikedImage", userID, imageURL).Return(nil).Run(func(args mock.Arguments) {
		b.likedImages[userID] = append(b.likedImages[userID], imageURL)
	})
	return b
}

// WithAddLikedImageError sets up the mock to handle AddLikedImage calls with an error (e.g., duplicate).
func (b *MockLikedImagesBuilder) WithAddLikedImageError(userID, imageURL string) *MockLikedImagesBuilder {
	b.mock.On("AddLikedImage", userID, imageURL).Return(e.NewError(e.ValidationErr, e.ImageAlreadyLiked, "image already liked", nil))
	return b
}

// WithLikedImages sets up the mock to return a list of liked images.
func (b *MockLikedImagesBuilder) WithGetLikedImages(userID string) *MockLikedImagesBuilder {
	images, exists := b.likedImages[userID]
	if !exists {
		images = []string{}
	}
	b.mock.On("GetLikedImages", userID).Return(images, nil)
	return b
}

// WithRemovedLikedImage sets up the mock to return a successful response when removing a liked image.
func (b *MockLikedImagesBuilder) WithRemoveLikedImage(userID, imageURL string) *MockLikedImagesBuilder {
	b.mock.On("RemoveLikedImage", userID, imageURL).Return(nil).Run(func(args mock.Arguments) {
		for i, img := range b.likedImages[userID] {
			if img == imageURL {
				b.likedImages[userID] = append(b.likedImages[userID][:i], b.likedImages[userID][i+1:]...)
				break
			}
		}
	})
	return b
}

func (b *MockLikedImagesBuilder) Build() *MockLikedImagesRepository {
	return b.mock
}

func (b *MockLikedImagesBuilder) AssertExpectations(t mock.TestingT) {
	b.mock.AssertExpectations(t)
}

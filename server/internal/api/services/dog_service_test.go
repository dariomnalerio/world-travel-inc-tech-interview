package services_test

import (
	s "server/internal/api/services"
	e "server/internal/errors"
	testing_mocks "server/internal/testing"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomImage(t *testing.T) {

	const (
		validImageURL = "https://dog.ceo/api/img/breed/image.jpg"
		emptyURL      = ""
	)

	t.Run("successful random image fetch", func(t *testing.T) {
		builder := testing_mocks.NewDogMockBuilder()
		builder.WithSuccessfulRandomPicture(validImageURL)
		likedImagesRepo := &testing_mocks.MockLikedImagesRepository{}
		service := s.NewDogService(builder.Build(), likedImagesRepo)
		response, err := service.GetRandomImage()

		assert.Equal(t, validImageURL, response)
		assert.NoError(t, err)
		builder.AssertExpectations(t)
	})

	t.Run("failed random image fetch", func(t *testing.T) {
		builder := testing_mocks.NewDogMockBuilder()
		likedImagesRepo := &testing_mocks.MockLikedImagesRepository{}
		error := e.NewError(e.InternalErr, e.ExternalAPIError, "failed to fetch random dog picture", nil)
		builder.WithFailedRandomPicture(error)
		service := s.NewDogService(builder.Build(), likedImagesRepo)

		_, err := service.GetRandomImage()

		assert.Error(t, err)
		assert.Equal(t, error, err)
		builder.AssertExpectations(t)
	})

	t.Run("empty image URL", func(t *testing.T) {
		builder := testing_mocks.NewDogMockBuilder()
		likedImagesRepo := &testing_mocks.MockLikedImagesRepository{}
		builder.WithSuccessfulRandomPicture(emptyURL)
		service := s.NewDogService(builder.Build(), likedImagesRepo)

		_, err := service.GetRandomImage()

		assert.Error(t, err)
		assert.Equal(t, e.NewError(e.ValidationErr, e.EmptyImageURL, "empty image URL", nil), err)
		builder.AssertExpectations(t)
	})

	t.Run("malformed URL - no protocol", func(t *testing.T) {
		likedImagesRepo := &testing_mocks.MockLikedImagesRepository{}
		builder := testing_mocks.NewDogMockBuilder()
		builder.WithSuccessfulRandomPicture("dog.ceo/image.jpg")
		service := s.NewDogService(builder.Build(), likedImagesRepo)

		_, err := service.GetRandomImage()

		assert.Error(t, err)
		assert.Equal(t, e.NewError(e.ValidationErr, e.MalformedURL, "malformed or invalid image URL", nil), err)
		builder.AssertExpectations(t)
	})

	t.Run("malformed URL - empty space in url", func(t *testing.T) {
		likedImagesRepo := &testing_mocks.MockLikedImagesRepository{}
		builder := testing_mocks.NewDogMockBuilder()
		builder.WithSuccessfulRandomPicture("https://dog.ceo/ima ge.jpg")
		service := s.NewDogService(builder.Build(), likedImagesRepo)

		_, err := service.GetRandomImage()

		assert.Error(t, err)
		assert.Equal(t, e.NewError(e.ValidationErr, e.MalformedURL, "URL contains empty spaces", nil), err)
		builder.AssertExpectations(t)
	})

	t.Run("malformed URL - incomplete URL", func(t *testing.T) {
		likedImagesRepo := &testing_mocks.MockLikedImagesRepository{}
		builder := testing_mocks.NewDogMockBuilder()
		builder.WithSuccessfulRandomPicture("https://")
		service := s.NewDogService(builder.Build(), likedImagesRepo)

		_, err := service.GetRandomImage()

		assert.Error(t, err)
		assert.Equal(t, e.NewError(e.ValidationErr, e.MalformedURL, "malformed or invalid image URL", nil), err)
		builder.AssertExpectations(t)
	})

	t.Run("malformed URL - invalid protocol", func(t *testing.T) {
		likedImagesRepo := &testing_mocks.MockLikedImagesRepository{}
		builder := testing_mocks.NewDogMockBuilder()
		builder.WithSuccessfulRandomPicture("ftp://dog.ceo/image.jpg")
		service := s.NewDogService(builder.Build(), likedImagesRepo)

		_, err := service.GetRandomImage()

		assert.Error(t, err)
		assert.Equal(t, e.NewError(e.ValidationErr, e.MalformedURL, "malformed or invalid image URL", nil), err)
		builder.AssertExpectations(t)
	})

	t.Run("malformed URL - no image URL", func(t *testing.T) {
		likedImagesRepo := &testing_mocks.MockLikedImagesRepository{}
		builder := testing_mocks.NewDogMockBuilder()
		builder.WithSuccessfulRandomPicture("https://dog.ceo/")
		service := s.NewDogService(builder.Build(), likedImagesRepo)

		_, err := service.GetRandomImage()

		assert.Error(t, err)
		assert.Equal(t, e.NewError(e.ValidationErr, e.InvalidImageExtension, "invalid image extension", nil), err)
		builder.AssertExpectations(t)
	})
}

package services_test

import (
	s "server/internal/api/services"
	testing_mocks "server/internal/testing"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	userID = "1"
)

const successImageURL = "https://example.com/image.jpg"

func TestGetLikedImages(t *testing.T) {
	t.Run("successful liked images retrieval", func(t *testing.T) {
		images := []string{"https://example.com/image1.jpg", "https://example.com/image2.jpg"}
		initialLikedImages := map[string][]string{
			userID: images,
		}
		userBuilder := testing_mocks.NewMockBuilder().WithFoundByID()
		likedImagesBuilder := testing_mocks.NewLikedImagesMockBuilder().WithInitialLikedImages(initialLikedImages).WithGetLikedImages(userID)
		service := s.NewLikedImagesService(likedImagesBuilder.Build(), userBuilder.Build())

		response, err := service.GetLikedImages(userID)

		assert.Equal(t, images, response)
		assert.NoError(t, err)
		userBuilder.AssertExpectations(t)
		likedImagesBuilder.AssertExpectations(t)
	})

	t.Run("empty image array retrieval", func(t *testing.T) {
		userBuilder := testing_mocks.NewMockBuilder().WithFoundByID()
		likedImagesBuilder := testing_mocks.NewLikedImagesMockBuilder().WithGetLikedImages(userID)
		service := s.NewLikedImagesService(likedImagesBuilder.Build(), userBuilder.Build())

		_, err := service.GetLikedImages(userID)

		assert.Equal(t, nil, err)
		userBuilder.AssertExpectations(t)
		likedImagesBuilder.AssertExpectations(t)
	})
}

func TestAddLikedImage(t *testing.T) {
	t.Run("successful liked image - returns void", func(t *testing.T) {
		userBuilder := testing_mocks.NewMockBuilder().WithFoundByID()
		likedImagesBuilder := testing_mocks.NewLikedImagesMockBuilder().WithGetLikedImages(userID).WithAddLikedImage(userID, successImageURL)
		service := s.NewLikedImagesService(likedImagesBuilder.Build(), userBuilder.Build())

		err := service.LikeImage(userID, successImageURL)
		assert.NoError(t, err)

		userBuilder.AssertExpectations(t)
		likedImagesBuilder.AssertExpectations(t)
	})

	t.Run("empty image URL - returns validation error", func(t *testing.T) {
		userBuilder := testing_mocks.NewMockBuilder().WithFoundByID()
		likedImagesBuilder := testing_mocks.NewLikedImagesMockBuilder()
		service := s.NewLikedImagesService(likedImagesBuilder.Build(), userBuilder.Build())

		err := service.LikeImage(userID, "")
		assert.Error(t, err)

		userBuilder.AssertExpectations(t)
		likedImagesBuilder.AssertExpectations(t)
	})

	t.Run("image was already liked - returns validation error", func(t *testing.T) {
		images := []string{successImageURL}
		initialLikedImages := map[string][]string{
			userID: images,
		}
		userBuilder := testing_mocks.NewMockBuilder().WithFoundByID()
		likedImagesBuilder := testing_mocks.NewLikedImagesMockBuilder().WithInitialLikedImages(initialLikedImages).WithGetLikedImages(userID)
		service := s.NewLikedImagesService(likedImagesBuilder.Build(), userBuilder.Build())

		err := service.LikeImage(userID, successImageURL)

		assert.Error(t, err)
		userBuilder.AssertExpectations(t)
		likedImagesBuilder.AssertExpectations(t)
	})

	t.Run("user not found - returns db error", func(t *testing.T) {
		userBuilder := testing_mocks.NewMockBuilder().WithErrorFindByID()
		likedImagesBuilder := testing_mocks.NewLikedImagesMockBuilder()
		service := s.NewLikedImagesService(likedImagesBuilder.Build(), userBuilder.Build())

		err := service.LikeImage(userID, successImageURL)

		assert.Error(t, err)
		userBuilder.AssertExpectations(t)
		likedImagesBuilder.AssertExpectations(t)
	})

}
func TestUnlikeImage(t *testing.T) {
	t.Run("successful unlike image - returns void", func(t *testing.T) {
		images := []string{successImageURL}
		initialLikedImages := map[string][]string{
			userID: images,
		}
		userBuilder := testing_mocks.NewMockBuilder().WithFoundByID()
		likedImagesBuilder := testing_mocks.NewLikedImagesMockBuilder().WithInitialLikedImages(initialLikedImages).WithGetLikedImages(userID).WithRemoveLikedImage(userID, successImageURL)
		service := s.NewLikedImagesService(likedImagesBuilder.Build(), userBuilder.Build())

		err := service.UnlikeImage(userID, successImageURL)
		assert.NoError(t, err)

		userBuilder.AssertExpectations(t)
		likedImagesBuilder.AssertExpectations(t)
	})

	t.Run("empty image URL - returns validation error", func(t *testing.T) {
		userBuilder := testing_mocks.NewMockBuilder().WithFoundByID()
		likedImagesBuilder := testing_mocks.NewLikedImagesMockBuilder()
		service := s.NewLikedImagesService(likedImagesBuilder.Build(), userBuilder.Build())

		err := service.UnlikeImage(userID, "")
		assert.Error(t, err)

		userBuilder.AssertExpectations(t)
		likedImagesBuilder.AssertExpectations(t)
	})

	t.Run("image not liked - returns validation error", func(t *testing.T) {
		images := []string{}
		initialLikedImages := map[string][]string{
			userID: images,
		}
		userBuilder := testing_mocks.NewMockBuilder().WithFoundByID()
		likedImagesBuilder := testing_mocks.NewLikedImagesMockBuilder().WithInitialLikedImages(initialLikedImages).WithGetLikedImages(userID)
		service := s.NewLikedImagesService(likedImagesBuilder.Build(), userBuilder.Build())

		err := service.UnlikeImage(userID, successImageURL)

		assert.Error(t, err)
		userBuilder.AssertExpectations(t)
		likedImagesBuilder.AssertExpectations(t)
	})

	t.Run("user not found - returns db error", func(t *testing.T) {
		userBuilder := testing_mocks.NewMockBuilder().WithErrorFindByID()
		likedImagesBuilder := testing_mocks.NewLikedImagesMockBuilder()
		service := s.NewLikedImagesService(likedImagesBuilder.Build(), userBuilder.Build())

		err := service.UnlikeImage(userID, successImageURL)

		assert.Error(t, err)
		userBuilder.AssertExpectations(t)
		likedImagesBuilder.AssertExpectations(t)
	})
}

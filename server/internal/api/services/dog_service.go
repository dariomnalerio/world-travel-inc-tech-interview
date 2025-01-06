package services

import (
	"server/internal/api/repositories"
	"server/internal/errors"
	"server/internal/utils"
)

type DogService struct {
	dogRepo       repositories.DogRepository
	likedImgsRepo repositories.LikedImagesRepository
}

func NewDogService(dogRepo repositories.DogRepository, likedImgsRepo repositories.LikedImagesRepository) *DogService {
	return &DogService{
		dogRepo:       dogRepo,
		likedImgsRepo: likedImgsRepo,
	}
}

func (s *DogService) GetRandomImage() (string, error) {
	imageURL, err := s.dogRepo.GetRandomPicture()

	if err != nil {
		return "", errors.NewError(errors.InternalErr, errors.ExternalAPIError, "failed to fetch random dog picture", nil)
	}

	if utils.IsEmptyString(imageURL) {
		return "", errors.NewError(errors.ValidationErr, errors.EmptyImageURL, "empty image URL", nil)
	}

	if utils.ContainsEmptySpace(imageURL) {
		return "", errors.NewError(errors.ValidationErr, errors.MalformedURL, "URL contains empty spaces", nil)
	}

	if utils.IsInvalidProtocol(imageURL) || utils.IsMalformedURL(imageURL) {
		return "", errors.NewError(errors.ValidationErr, errors.MalformedURL, "malformed or invalid image URL", nil)
	}

	if !utils.HasImageValidExtension(imageURL) {
		return "", errors.NewError(errors.ValidationErr, errors.InvalidImageExtension, "invalid image extension", nil)
	}

	return imageURL, nil
}

// GetRandomImageAndCheckLike returns a random dog image URL and checks if the image has been liked by the user.
// It takes a user ID as input and returns the image URL, a boolean indicating if the image has been liked by the user,
// and an error if any.
func (s *DogService) GetRandomImageAndCheckLike(userID string) (string, bool, error) {
	imageURL, err := s.GetRandomImage()
	if err != nil {
		return "", false, err
	}

	likedImages, err := s.likedImgsRepo.GetLikedImages(userID)
	if err != nil {
		return "", false, err
	}

	for _, img := range likedImages {
		if img == imageURL {
			return img, true, nil
		}
	}

	return imageURL, false, nil
}

package services

import (
	"server/internal/api/repositories"
	"server/internal/errors"
	"server/internal/utils"
)

type DogService struct {
	dogRepo repositories.DogRepository
}

func NewDogService(dogRepo repositories.DogRepository) *DogService {
	return &DogService{
		dogRepo: dogRepo,
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

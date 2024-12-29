package services

import (
	"server/internal/api/repositories"
	e "server/internal/errors"
	"server/internal/utils"
)

type LikedImagesService struct {
	likedRepo repositories.LikedImagesRepository
	userRepo  repositories.UserRepository
}

func NewLikedImagesService(likedRepo repositories.LikedImagesRepository, userRepo repositories.UserRepository) *LikedImagesService {
	return &LikedImagesService{
		likedRepo,
		userRepo,
	}
}

func (s *LikedImagesService) GetLikedImages(userID string) ([]string, error) {
	_, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, e.NewError(e.UserErr, e.UserNotFound, "user not found", err)
	}

	return s.likedRepo.GetLikedImages(userID)
}

func (s *LikedImagesService) LikeImage(userID, imageURL string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil || user == nil {
		return e.NewError(e.UserErr, e.UserNotFound, "user not found", err)
	}

	if utils.IsEmptyString(imageURL) {
		return e.NewError(e.ValidationErr, e.EmptyImageURL, "empty image URL", nil)
	}

	if utils.ContainsEmptySpace(imageURL) {
		return e.NewError(e.ValidationErr, e.MalformedURL, "URL contains empty spaces", nil)
	}

	if utils.IsInvalidProtocol(imageURL) || utils.IsMalformedURL(imageURL) {
		return e.NewError(e.ValidationErr, e.MalformedURL, "malformed or invalid image URL", nil)
	}

	if !utils.HasImageValidExtension(imageURL) {
		return e.NewError(e.ValidationErr, e.InvalidImageExtension, "invalid image extension", nil)
	}

	imgs, err := s.likedRepo.GetLikedImages(userID)

	if err != nil {
		return e.NewError(e.InternalErr, e.DatabaseError, "error retrieving liked images", err)
	}

	if utils.IsImageLiked(imgs, imageURL) {
		return e.NewError(e.ValidationErr, e.ImageAlreadyLiked, "image already liked", nil)
	}

	return s.likedRepo.AddLikedImage(userID, imageURL)
}

func (s *LikedImagesService) UnlikeImage(userID, imageURL string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil || user == nil {
		return e.NewError(e.UserErr, e.UserNotFound, "user not found", err)
	}

	if utils.IsEmptyString(imageURL) {
		return e.NewError(e.ValidationErr, e.EmptyImageURL, "empty image URL", nil)
	}

	if utils.ContainsEmptySpace(imageURL) {
		return e.NewError(e.ValidationErr, e.MalformedURL, "URL contains empty spaces", nil)
	}

	if utils.IsInvalidProtocol(imageURL) || utils.IsMalformedURL(imageURL) {
		return e.NewError(e.ValidationErr, e.MalformedURL, "malformed or invalid image URL", nil)
	}

	if !utils.HasImageValidExtension(imageURL) {
		return e.NewError(e.ValidationErr, e.InvalidImageExtension, "invalid image extension", nil)
	}

	imgs, err := s.likedRepo.GetLikedImages(userID)

	if err != nil {
		return e.NewError(e.InternalErr, e.DatabaseError, "error retrieving liked images", err)
	}

	if !utils.IsImageLiked(imgs, imageURL) {
		return e.NewError(e.ValidationErr, e.ImageNotLiked, "image not liked", nil)
	}

	return s.likedRepo.RemoveLikedImage(userID, imageURL)
}

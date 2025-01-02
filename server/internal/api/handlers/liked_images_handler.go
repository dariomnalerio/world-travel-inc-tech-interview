package handlers

import (
	"log"
	"net/http"
	"server/internal/api/services"
	"server/internal/models"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type LikedImagesHandler struct {
	likedImagesService *services.LikedImagesService
}

func NewLikedImagesHandler(likedImagesService *services.LikedImagesService) *LikedImagesHandler {
	return &LikedImagesHandler{
		likedImagesService: likedImagesService,
	}
}

// GetLikedImages godoc
//
//	@Summary		Returns a list of liked images.
//	@Description	Returns a list of liked images for the user.
//	@Tags			liked_images
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Success		200		{object}	models.GetLikedImagesResponse
//	@Failure		400		{object}	utils.ErrorResponse
//
//	@Security		BearerAuth
//
//	@Router			/liked_images/{id} [get]
func (h *LikedImagesHandler) GetLikedImages(c *gin.Context) {
	log.Default().Println("Getting liked images")
	var req models.GetLikedImagesRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	imgs, err := h.likedImagesService.GetLikedImages(req.UserID)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"images": imgs,
	})
}

// LikeImage godoc
//
//	@Summary		Likes an image.
//	@Description	Likes an image for the user.
//	@Tags			liked_images
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Param			request body models.LikeImageRequestBody true "Image URL"
//	@Success		201		{object}	models.LikeImageResponse
//	@Failure		400		{object}	utils.ErrorResponse
//
//	@Security		BearerAuth
//
//	@Router			/liked_images/{id} [post]
func (h *LikedImagesHandler) LikeImage(c *gin.Context) {
	var body models.LikeImageRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var req models.LikeImageRequestURL
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err := h.likedImagesService.LikeImage(req.UserID, body.ImageURL)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": "true",
		"image":   body.ImageURL,
	})
}

// UnlikeImage godoc
//
//	@Summary		Unlikes an image.
//	@Description	Unlikes an image for the user.
//	@Tags			liked_images
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Param			request body models.UnlikeImageRequestBody true "Image URL"
//	@Success		200		{object}	models.UnlikeImageResponse
//	@Failure		400		{object}	utils.ErrorResponse
//
//	@Security		BearerAuth
//
//	@Router			/liked_images/{id} [delete]
func (h *LikedImagesHandler) UnlikeImage(c *gin.Context) {

	var body models.UnlikeImageRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var req models.UnlikeImageRequestURL
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID in URI"})
		return
	}

	err := h.likedImagesService.UnlikeImage(req.UserID, body.ImageURL)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "true",
		"image":   body.ImageURL,
	})
}

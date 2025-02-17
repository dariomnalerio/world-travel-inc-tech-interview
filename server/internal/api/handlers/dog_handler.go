package handlers

import (
	"net/http"
	"server/internal/api/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type DogHandler struct {
	dogHandler *services.DogService
}

func NewDogHandler(dogService *services.DogService) *DogHandler {
	return &DogHandler{
		dogService,
	}
}

// GetRandomImage godoc
//
//	@Summary		Returns a random dog image URL.
//	@Description	Returns a random dog image URL from the Dog API.
//	@Tags			dog
//	@Accept			json
//	@Produce		json
//
//	@Param			userID	query	string	false	"User ID"
//
//	@Success		200		{string}	models.GetRandomImageResponse
//	@Failure		400		{object}	utils.ErrorResponse
//	@Router			/dog/random [get]
func (h *DogHandler) GetRandomImage(c *gin.Context) {
	userID := c.DefaultQuery("userID", "")

	if userID == "" {
		img, err := h.dogHandler.GetRandomImage()
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"image_url": img,
			"liked":     false,
		})
		return
	}

	img, isLiked, err := h.dogHandler.GetRandomImageAndCheckLike(userID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"image_url": img,
		"liked":     isLiked,
	})
}

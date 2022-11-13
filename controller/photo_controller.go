package controller

import (
	"net/http"
	"project2/helper"
	"project2/model/input"
	"project2/model/response"
	"project2/service"

	"github.com/gin-gonic/gin"
)

type photoController struct {
	photoService service.PhotoService
	userService  service.UserService
}

func NewPhotoController(photoService service.PhotoService, userService service.UserService) *photoController {
	return &photoController{photoService, userService}
}

func (h *photoController) AddNewPhoto(c *gin.Context) {
	var input input.PhotoCreateInput

	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// send to service
	newPhoto, err := h.photoService.CreatePhoto(input, currentUser)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPhotoResponse := response.CreatePhotoResponse{
		ID:       newPhoto.ID,
		Title:    newPhoto.Title,
		Caption:  newPhoto.Caption,
		PhotoURL: input.PhotoURL,
		UserID:   currentUser,
	}

	response := helper.APIResponse("created", newPhotoResponse)
	c.JSON(http.StatusOK, response)
}

func (h *photoController) DeletePhoto(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	var idPhotoUri input.PhotoDeleteIDUser

	err := c.ShouldBindUri(&idPhotoUri)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	idPhoto := idPhotoUri.ID

	if idPhoto == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = h.photoService.DeletePhoto(idPhoto, currentUser)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deleteResponse := response.PhotoDeleteResponse{
		Message: "Your photo has been successfully deleted",
	}

	response := helper.APIResponse("ok", deleteResponse)
	c.JSON(http.StatusOK, response)
}

func (h *photoController) GetPhotos(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	photos, err := h.photoService.GetPhotosAll()
	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var photoResponse []response.GetPhotoUser

	for _, index := range photos {

		userTmp, _ := h.userService.GetUserByID(index.UserID)

		photoResponseTmp := response.GetPhotoUser{
			ID:        index.ID,
			Title:     index.Title,
			Caption:   index.Caption,
			PhotoURL:  index.PhotoURL,
			UserID:    index.UserID,
			CreatedAt: index.CreatedAt,
			UpdatedAt: index.UpdatedAt,
			User: response.UserInPhoto{
				Username: userTmp.Username,
				Email:    userTmp.Email,
			},
		}

		photoResponse = append(photoResponse, photoResponseTmp)
	}

	response := helper.APIResponse("ok", photoResponse)
	c.JSON(http.StatusOK, response)
}

func (h *photoController) GetPhoto(c *gin.Context) {

	var idPhotoUri input.PhotoDeleteIDUser

	err := c.ShouldBindUri(&idPhotoUri)
	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	idPhoto := idPhotoUri.ID

	if idPhoto == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	photo, err := h.photoService.GetPhotoByID(idPhoto)

	if err != nil {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, err := h.userService.GetUserByID(photo.UserID)
	if err != nil {
		response := helper.APIResponse("failed", "user not found!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// _, _ := h.commentService.GetCommentsByPhotoID(idPhoto)

	photoResponse := response.GetPhotoUser{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
		User: response.UserInPhoto{
			Username: user.Username,
			Email:    user.Email,
		},
	}

	response := helper.APIResponse("ok", photoResponse)
	c.JSON(http.StatusOK, response)
}

func (h *photoController) UpdatePhoto(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	updatePhoto := input.PhotoUpdateInput{}

	err := c.ShouldBindJSON(&updatePhoto)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var idPhotoUri input.PhotoUpdateIDUser

	err = c.ShouldBindUri(&idPhotoUri)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	id_photo := idPhotoUri.ID

	_, err = h.photoService.UpdatePhoto(currentUser, id_photo, updatePhoto)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	photoUpdated, _ := h.photoService.GetPhotoByID(id_photo)

	photoResponse := response.PhotoUpdateResponse{
		ID:        photoUpdated.ID,
		Title:     photoUpdated.Title,
		Caption:   photoUpdated.Caption,
		PhotoURL:  photoUpdated.PhotoURL,
		UserID:    photoUpdated.UserID,
		UpdatedAt: photoUpdated.UpdatedAt,
	}

	response := helper.APIResponse("ok", photoResponse)
	c.JSON(http.StatusOK, response)
}

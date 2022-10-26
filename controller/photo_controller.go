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
	photoService   service.PhotoService
	commentService service.CommentService
	userSerice     service.UserService
}

func NewPhotoController(photoService service.PhotoService, commentService service.CommentService, userService service.UserService) *photoController {
	return &photoController{photoService, commentService, userService}
}

func (h *photoController) AddNewPhoto(c *gin.Context) {
	var input input.InputPhotos

	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send to service
	newPhoto, err := h.photoService.CreatePhoto(input, currentUser)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)

		response := helper.APIResponse("failed", errorMessages)
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
	return
}

func (h *photoController) DeletePhoto(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	var idPhotoUri input.DeletePhoto

	err := c.ShouldBindUri(&idPhotoUri)

	if err != nil {
		response := helper.APIResponse("failed", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	idPhoto := idPhotoUri.ID

	if idPhoto == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = h.photoService.DeletePhoto(idPhoto)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("ok", "success deleted photo!")
	c.JSON(http.StatusOK, response)
	return
}

func (h *photoController) GetPhotos(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	photos, err := h.photoService.GetPhotosUser(currentUser)

	var photoResponse []response.GetPhotoUser

	for _, index := range photos {

		commentTmp, _ := h.commentService.GetCommentsByPhotoID(index.ID)

		photoResponseTmp := response.GetPhotoUser{
			ID:        index.ID,
			Title:     index.Title,
			Caption:   index.Caption,
			PhotoURL:  index.PhotoURL,
			CreatedAt: index.CreatedAt,
			Comments:  commentTmp,
		}

		photoResponse = append(photoResponse, photoResponseTmp)
	}

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("ok", photoResponse)
	c.JSON(http.StatusOK, response)
}

func (h *photoController) GetPhoto(c *gin.Context) {

	var idPhotoUri input.DeletePhoto

	err := c.ShouldBindUri(&idPhotoUri)

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

	user, err := h.userSerice.GetUserByID(photo.UserID)
	if err != nil {
		response := helper.APIResponse("failed", "user not found!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	comments, _ := h.commentService.GetCommentsByPhotoID(idPhoto)

	photoResponse := response.GetPhotoDetailUser{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		CreatedAt: photo.CreatedAt,
		Comments:  comments,
		User: response.UserInPhoto{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}

	response := helper.APIResponse("ok", photoResponse)
	c.JSON(http.StatusOK, response)
	return
}

func (h *photoController) UpdatePhoto(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	updatePhoto := input.UpdatePhoto{}

	err := c.ShouldBindJSON(&updatePhoto)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var idPhotoUri input.UpdatePhotoIDUser

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

	_, err = h.photoService.UpdatePhoto(id_photo, updatePhoto)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	photoUpdated, _ := h.photoService.GetPhotoByID(id_photo)

	response := helper.APIResponse("ok", photoUpdated)
	c.JSON(http.StatusOK, response)
}

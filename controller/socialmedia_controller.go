package controller

import (
	"net/http"
	"project2/helper"
	"project2/model/input"
	"project2/model/response"
	"project2/service"

	"github.com/gin-gonic/gin"
)

type socialmediaController struct {
	socialmediaService service.SocialMediaService
	userService        service.UserService
}

func NewSocialMediaController(socialmediaService service.SocialMediaService, userService service.UserService) *socialmediaController {
	return &socialmediaController{socialmediaService, userService}
}

func (h *socialmediaController) AddNewSocialMedia(c *gin.Context) {
	var input input.SocialInput

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
	newSocialMedia, err := h.socialmediaService.CreateSocialMedia(input, currentUser)

	if err != nil {
		// errorMessages := helper.FormatValidationError(err)

		response := helper.APIResponse("failed", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSocialMediaResponse := response.SocialMediaCreateResponse{
		ID:        newSocialMedia.ID,
		Name:      newSocialMedia.Name,
		URL:       newSocialMedia.URL,
		UsedID:    newSocialMedia.UserID,
		CreatedAt: newSocialMedia.CreatedAt,
	}

	response := helper.APIResponse("created", newSocialMediaResponse)
	c.JSON(http.StatusCreated, response)
}

func (h *socialmediaController) DeleteSocialmedia(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	var idSocialMediaUri input.DeleteSocialMedia

	err := c.ShouldBindUri(&idSocialMediaUri)

	if err != nil {
		response := helper.APIResponse("failed", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	idSocialMedia := idSocialMediaUri.ID

	if idSocialMedia == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = h.socialmediaService.DeleteSocialMedia(idSocialMedia)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("ok", "success deleted social media!")
	c.JSON(http.StatusOK, response)
	return
}

func (h *socialmediaController) GetSocialMedia(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	socialmedia, err := h.socialmediaService.GetSocialMedia(currentUser)
	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userService.GetUserByID(currentUser)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	responseSocialMedia, err := response.GetAllSocialMedia(socialmedia, user)
	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("ok", responseSocialMedia)
	c.JSON(http.StatusOK, response)
}

func (h *socialmediaController) UpdateSocialMedia(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	update := input.SocialInput{}

	err := c.ShouldBindJSON(&update)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var idSocialUri input.UpdateSocialMedia

	err = c.ShouldBindUri(&idSocialUri)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	id_socialmedia := idSocialUri.ID

	queryResult, err := h.socialmediaService.UpdateSocialMedia(id_socialmedia, update)

	if queryResult.ID == 0 {
		response := helper.APIResponse("failed", "photo not found!")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	Updated, err := h.socialmediaService.GetSocialMedia(id_socialmedia)

	response := helper.APIResponse("ok", Updated)
	c.JSON(http.StatusOK, response)
	return
}

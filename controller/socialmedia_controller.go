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

// Add New Social Media
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

	// Send To Service
	newSocialMedia, err := h.socialmediaService.CreateSocialMedia(input, currentUser)

	if err != nil {

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

// Delete Social Media by id
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
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	idSocialMedia := idSocialMediaUri.ID

	if idSocialMedia == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = h.socialmediaService.DeleteSocialMedia(currentUser, idSocialMedia)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	responseSocialMedia := response.SocialMediaDeleteResponse{
		Message: "Your social media has been successfully deleted",
	}

	response := helper.APIResponse("ok", responseSocialMedia)
	c.JSON(http.StatusOK, response)
}

// Get All Social Medias
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

// Edit/Update Social Media (Name or URL)
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

	_, err = h.socialmediaService.UpdateSocialMedia(currentUser, id_socialmedia, update)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	socialMediaUpdated, err := h.socialmediaService.GetSocialMediaByID(id_socialmedia)
	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	responseSocialMedia := response.SocialMediaUpdateResponse{
		ID:        socialMediaUpdated.ID,
		Name:      socialMediaUpdated.Name,
		URL:       socialMediaUpdated.URL,
		UserID:    socialMediaUpdated.UserID,
		UpdatedAt: socialMediaUpdated.UpdatedAt,
	}

	response := helper.APIResponse("ok", responseSocialMedia)
	c.JSON(http.StatusOK, response)
}

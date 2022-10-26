package controller

import (
	"net/http"
	"project2/helper"
	"project2/model/input"
	"project2/model/response"
	"project2/service"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentService service.CommentService
	photoService   service.PhotoService
}

func NewCommentController(commentService service.CommentService, photoService service.PhotoService) *commentController {
	return &commentController{commentService, photoService}
}

func (h *commentController) AddNewComment(c *gin.Context) {
	var input input.CommentInput

	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorsMessage := helper.FormatValidationError(err)

		response := helper.APIResponse("failed", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send to service
	newComment, err := h.commentService.CreateComment(input, currentUser)

	if err != nil {

		response := helper.APIResponse("failed", err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCommentResponse := response.CreateCommentResponse{
		ID:      newComment.ID,
		Message: newComment.Message,
		PhotoID: input.PhotoID,
		UserID:  currentUser,
	}

	response := helper.APIResponse("created", newCommentResponse)
	c.JSON(http.StatusOK, response)
}

func (h *commentController) DeleteComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	var idCommentUri input.DeleteComment

	err := c.ShouldBindUri(&idCommentUri)

	if err != nil {
		response := helper.APIResponse("failed", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	idComment := idCommentUri.ID

	if idComment == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = h.commentService.DeleteComment(idComment)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("ok", "success deleted comment!")
	c.JSON(http.StatusOK, response)
}

func (h *commentController) GetComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	// var idPhotoUri input.UpdatePhotoIDUser

	// err := c.ShouldBindUri(&idPhotoUri)

	// if err != nil {
	// 	errorMessages := helper.FormatValidationError(err)
	// 	response := helper.APIResponse("failed", gin.H{
	// 		"errors": errorMessages,
	// 	})
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
	// 	return
	// }

	comments, _ := h.commentService.GetComment(currentUser)

	// query photo
	var allCommentsPhoto []response.GetCommentResponse
	for _, item := range comments {
		photo, _ := h.photoService.GetPhotoByID(item.PhotoID)
		allCommentsPhotoTmp := response.GetAllComment(item, photo)

		allCommentsPhoto = append(allCommentsPhoto, allCommentsPhotoTmp)
	}

	response := helper.APIResponse("ok", allCommentsPhoto)
	c.JSON(http.StatusOK, response)

	// if err != nil {
	// 	errorMessages := helper.FormatValidationError(err)
	// 	response := helper.APIResponse("failed", gin.H{
	// 		"errors": errorMessages,
	// 	})
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// }

	// response := helper.APIResponse("ok", response.GetAllComment(comment, photo))
	// c.JSON(http.StatusOK, response)
}

func (h *commentController) UpdateComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	UpdateComment := input.CommentUpdateInput{}

	err := c.ShouldBindJSON(&UpdateComment)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var idCommentUri input.UpdateComment

	err = c.ShouldBindUri(&idCommentUri)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	id_comment := idCommentUri.ID

	_, err = h.commentService.UpdateComment(id_comment, UpdateComment)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	Updated, _ := h.commentService.GetCommentByID(id_comment)

	if Updated.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, "comment not found")
		return
	}

	response := helper.APIResponse("ok", Updated)
	c.JSON(http.StatusOK, response)
}

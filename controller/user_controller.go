package controller

import (
	"net/http"
	"project2/helper"
	"project2/middleware"
	"project2/model/input"
	"project2/model/response"
	"project2/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input input.RegisterUserInput

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
	newUser, err := h.userService.CreateUser(input)

	if err != nil {
		response := helper.APIResponse("failed", err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUserResponse := response.UserRegisterResponse{
		ID:        newUser.ID,
		Age:       newUser.Age,
		Email:     newUser.Email,
		Password:  newUser.Password,
		Username:  newUser.Username,
		CreatedAt: newUser.CreatedAt,
	}

	response := helper.APIResponse("created", newUserResponse)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userController) Login(c *gin.Context) {
	var input input.LoginUserInput

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

	// send to services
	// get user by email
	user, err := h.userService.GetUserByEmail(input.Email)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// return when user not found!
	if user.ID == 0 {
		errorMessages := "User not found!"
		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusNotFound, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		response := helper.APIResponse("failed", "password not match!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// lets create token!
	jwtService := middleware.NewService()
	token, err := jwtService.GenerateToken(user.ID)

	// return the token!
	response := helper.APIResponse("ok", gin.H{
		"token": token,
	})
	c.JSON(http.StatusOK, response)
	return
}

func (h *userController) UpdateUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputUserUpdate input.UpdateUserInput

	err := c.ShouldBindJSON(&inputUserUpdate)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	_, err = h.userService.UpdateUser(currentUser, inputUserUpdate)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	userUpdated, err := h.userService.GetUserByID(currentUser)

	response := helper.APIResponse("ok", userUpdated)
	c.JSON(http.StatusOK, response)
	return
}

func (h *userController) DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	_, err := h.userService.DeleteUser(currentUser)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("ok", "success deleted user!")
	c.JSON(http.StatusOK, response)
	return
}

func (h *userController) TestUser(c *gin.Context) {
	id_user, err := c.Get("currentUser")

	if err == false {
		c.JSON(http.StatusNotFound, helper.APIResponse("not created", err))
		return
	}

	c.JSON(http.StatusOK, helper.APIResponse("created", id_user))
	return
}

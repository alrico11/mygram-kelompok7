package controller

import (
	"fmt"
	"net/http"
	"project2/helper"
	"project2/middleware"
	"project2/model/input"
	"project2/model/response"
	"project2/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input input.UserRegisterInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		// log.Fatal(err)
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error On Filled %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			fmt.Println(err)
			return
		}

	}
	user, err := h.userService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": err.Error(),
		})
		return
	}

	registerResponse := response.UserRegisterResponse{
		ID:       user.ID,
		Age:      user.Age,
		Email:    user.Email,
		Username: user.Username,
	}

	response := helper.APIResponse("created", registerResponse)
	c.JSON(201, response)
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": convertToUserResponse(user),
	// })

	// // send to service
	// newUser, err := h.userService.CreateUser(input)

	// if err != nil {
	// 	response := helper.APIResponse("failed", err)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	// newUserResponse := response.UserRegisterResponse{
	// 	ID:        newUser.ID,
	// 	Age:       newUser.Age,
	// 	Email:     newUser.Email,
	// 	Password:  newUser.Password,
	// 	Username:  newUser.Username,
	// 	CreatedAt: newUser.CreatedAt,
	// }

	// response := helper.APIResponse("created", newUserResponse)
	// c.JSON(http.StatusOK, response)
	// return
}

func (h *userController) Login(c *gin.Context) {
	var input input.UserLoginInput

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
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
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

	// create token
	jwtService := middleware.NewService()
	token, err := jwtService.GenerateToken(user.ID)
	if err != nil {
		response := helper.APIResponse("failed", "failed to generate token!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginResponse := response.UserLoginResponse{
		Token: token,
	}

	// return token
	response := helper.APIResponse("ok", loginResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userController) UpdateUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	var inputUserUpdate input.UserUpdateInput

	err := c.ShouldBindJSON(&inputUserUpdate)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.userService.UpdateUser(currentUser, inputUserUpdate)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userUpdated, err := h.userService.GetUserByID(currentUser)
	if err != nil {
		response := helper.APIResponse("failed", "Cannot fetch user!")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updateResponse := response.UserUpdateResponse{
		ID:        userUpdated.ID,
		Email:     userUpdated.Email,
		Username:  userUpdated.Username,
		Age:       userUpdated.Age,
		UpdatedAt: userUpdated.UpdatedAt,
	}

	response := helper.APIResponse("ok", updateResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userController) DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	_, err := h.userService.DeleteUser(currentUser)

	if err != nil {
		response := helper.APIResponse("failed", gin.H{
			"errors": err.Error(),
		})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	deleteResponse := response.UserDeleteResponse{
		Message: "Your account has been successfully deleted",
	}

	response := helper.APIResponse("ok", deleteResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userController) TestUser(c *gin.Context) {
	id_user, err := c.Get("currentUser")

	if !err {
		c.JSON(http.StatusNotFound, helper.APIResponse("not created", err))
		return
	}

	c.JSON(http.StatusOK, helper.APIResponse("created", id_user))
}

// func convertToUserResponse(u entity.User) response.UserRegisterResponse {
// 	return response.UserRegisterResponse{
// 		Age:      u.Age,
// 		Email:    u.Email,
// 		ID:       u.ID,
// 		Password: u.Password,
// 		Username: u.Username,
// 	}
// }

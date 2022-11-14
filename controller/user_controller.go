package controller

import (
	"fmt"
	"net/http"
	"project2/helper"
	"project2/middleware"
	"project2/model/entity"
	"project2/model/input"
	"project2/model/response"
	"project2/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

func (h *userController) RegisterUser(c *gin.Context) {
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		// log.Fatal(err)
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error On Filled %s, condition: %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
			fmt.Println(err)
			return
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send to service
	newUser, err := h.userService.CreateUser(input)

	user, err := h.userService.CreateUser(input)
	if err != nil {
		response := helper.APIResponse("failed", err)
		c.JSON(http.StatusUnprocessableEntity, response)
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convertToUserResponse(user),
	})

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
@@ -106,11 +119,11 @@ func (h *userController) Login(c *gin.Context) {
		return
	}

	// lets create token!
	// crete token
	jwtService := middleware.NewService()
	token, err := jwtService.GenerateToken(user.ID)

	// return the token!
	// return token
	response := helper.APIResponse("ok", gin.H{
		"token": token,
	})
func (h *userController) DeleteUser(c *gin.Context) {
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("ok", "success deleted user!")
	response := helper.APIResponse("ok", "Success deleted user!")
	c.JSON(http.StatusOK, response)
	return
}
func (h *userController) TestUser(c *gin.Context) {

	c.JSON(http.StatusOK, helper.APIResponse("created", id_user))
	return

}

func convertToUserResponse(u entity.User) response.UserRegisterResponse {
	return response.UserRegisterResponse{
		Age:      u.Age,
		Email:    u.Email,
		Password: u.Password,
		Username: u.Username,
	}
}

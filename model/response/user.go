package response

import (
	"time"
    	"github.com/jinzhu/gorm"
)

type UserRegisterResponse struct {
	ID       int    `json:"id"gorm:"primary_key;auto_increment;not_null"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdateResponse struct {
	ID        int       `json:"id"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDeleteResponse struct {
	Message string `json:"message"`
}

package response

import (
	"time"
)

type UserRegisterResponse struct {
	ID       int `gorm:"primaryKey;autoIncrement;unique"`
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

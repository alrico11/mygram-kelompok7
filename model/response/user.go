package response

import "time"

type UserRegisterResponse struct {
	ID        int       `json:"id"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"date"`
}

type UserUpdateResponse struct {
	ID       int    `json:"id"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

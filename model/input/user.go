package input

type UserRegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"required,min=8"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateInput struct {
	Email    string `json:"email" binding:"email"`
	Username string `json:"username"`
	Password string `json:"password" binding:"omitempty,min=6"`
}

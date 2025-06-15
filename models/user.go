package models

type RegisterParameter struct {
	Name            string `json:"name" binding:"required,min=8,max=255"`
	Email           string `json:"email" binding:"required,email,max=255"`
	Password        string `json:"password" binding:"required,min=8,max=15"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=8,max=15,eqfield=Password"`
}

type LoginParameter struct {
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=8,max=15"`
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

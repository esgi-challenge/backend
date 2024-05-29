package models

type AuthLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthRegister struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

type Auth struct {
	Token string `json:"token" binding:"required"`
}

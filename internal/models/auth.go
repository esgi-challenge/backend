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

type AuthInvitationCode struct {
	InvitationCode string `json:"invitationCode" binding:"required"`
	Password       string `json:"password" binding:"required"`
}

type AuthResetPassword struct {
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Auth struct {
	Token string `json:"token" binding:"required"`
}

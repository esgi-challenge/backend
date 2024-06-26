package auth

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Login() gin.HandlerFunc
	Register() gin.HandlerFunc
	InvitationCode() gin.HandlerFunc
	ResetPassword() gin.HandlerFunc
}

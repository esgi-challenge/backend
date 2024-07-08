package user

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	GetAll() gin.HandlerFunc
	GetMe() gin.HandlerFunc
	SendResetMail() gin.HandlerFunc
	Create() gin.HandlerFunc
	UpdateMe() gin.HandlerFunc
	UpdateMePassword() gin.HandlerFunc
}

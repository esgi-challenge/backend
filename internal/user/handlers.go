package user

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	GetAll() gin.HandlerFunc
	SendResetMail() gin.HandlerFunc
  Create() gin.HandlerFunc
}

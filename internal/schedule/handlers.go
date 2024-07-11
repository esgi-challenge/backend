package schedule

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetById() gin.HandlerFunc
	GetSignatureCode() gin.HandlerFunc
	Sign() gin.HandlerFunc
	CheckSign() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

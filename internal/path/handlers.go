package path

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetById() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}
package school

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetById() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

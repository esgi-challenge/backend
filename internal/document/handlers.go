package document

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create() gin.HandlerFunc
	GetById() gin.HandlerFunc
  GetAllByUserId() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

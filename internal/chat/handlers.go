package chat

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetById() gin.HandlerFunc
  GetAllStudentChatter() gin.HandlerFunc
  GetAllTeacherChatter() gin.HandlerFunc
}

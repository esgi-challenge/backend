package school

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create() gin.HandlerFunc
	Invite() gin.HandlerFunc
	GetByUser() gin.HandlerFunc
	GetById() gin.HandlerFunc
	Delete() gin.HandlerFunc
	GetSchoolStudents() gin.HandlerFunc
	RemoveStudent() gin.HandlerFunc
	AddUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
}

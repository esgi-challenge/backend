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
	GetSchoolUsers() gin.HandlerFunc
	RemoveUser() gin.HandlerFunc
	AddUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
}

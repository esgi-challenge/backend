package class

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create() gin.HandlerFunc
	Add() gin.HandlerFunc
	Remove() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	GetById() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	GetClassLessStudents() gin.HandlerFunc
}

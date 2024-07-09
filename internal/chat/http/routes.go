package http

import (
	"github.com/esgi-challenge/backend/internal/chat"
	"github.com/gin-gonic/gin"
)

func SetupChatRoutes(chatGroup *gin.RouterGroup, h chat.Handlers) {
	chatGroup.POST("/channel", h.Create())
	chatGroup.GET("/channel", h.GetAll())
	chatGroup.GET("/students", h.GetAllStudentChatter())
	chatGroup.GET("/teachers", h.GetAllTeacherChatter())
	chatGroup.GET("/channel/:id", h.GetById())
}

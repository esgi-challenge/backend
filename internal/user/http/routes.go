package http

import (
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(userGroup *gin.RouterGroup, h user.Handlers) {
	userGroup.GET("", h.GetAll())
	userGroup.POST("/reset-link", h.SendResetMail())
}

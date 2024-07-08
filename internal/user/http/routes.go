package http

import (
	"github.com/esgi-challenge/backend/internal/user"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(userGroup *gin.RouterGroup, h user.Handlers) {
	userGroup.GET("", h.GetAll())
	userGroup.GET("/me", h.GetMe())
	userGroup.PUT("/me", h.UpdateMe())
	userGroup.PUT("/me/password", h.UpdateMePassword())
	userGroup.POST("", h.Create())
	userGroup.POST("/reset-link", h.SendResetMail())
}

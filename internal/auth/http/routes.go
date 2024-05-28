package http

import (
	"github.com/esgi-challenge/backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(schoolGroup *gin.RouterGroup, h auth.Handlers) {
	schoolGroup.POST("/login", h.Login())
	schoolGroup.POST("/register", h.Register())
}

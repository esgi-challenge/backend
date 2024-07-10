package http

import (
	"github.com/esgi-challenge/backend/internal/document"
	"github.com/gin-gonic/gin"
)

func SetupDocumentRoutes(documentGroup *gin.RouterGroup, h document.Handlers) {
	documentGroup.POST("", h.Create())
	documentGroup.GET("/:id", h.GetById())
	documentGroup.DELETE("/:id", h.Delete())
}

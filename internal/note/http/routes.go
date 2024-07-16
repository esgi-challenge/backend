package http

import (
	"github.com/esgi-challenge/backend/internal/note"
	"github.com/gin-gonic/gin"
)

func SetupNoteRoutes(noteGroup *gin.RouterGroup, h note.Handlers) {
	noteGroup.POST("", h.Create())
	noteGroup.GET("", h.GetAll())
	noteGroup.DELETE("/:id", h.Delete())
	noteGroup.PUT("/:id", h.Update())
}

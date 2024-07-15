package http

import (
	"github.com/esgi-challenge/backend/internal/informations"
	"github.com/gin-gonic/gin"
)

func SetupInformationsRoutes(informationsGroup *gin.RouterGroup, h informations.Handlers) {
	informationsGroup.POST("", h.Create())
	informationsGroup.GET("", h.GetAll())
	informationsGroup.GET("/:id", h.GetById())
	informationsGroup.DELETE("/:id", h.Delete())
}

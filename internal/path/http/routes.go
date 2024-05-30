package http

import (
	"github.com/esgi-challenge/backend/internal/path"
	"github.com/gin-gonic/gin"
)

func SetupPathRoutes(pathGroup *gin.RouterGroup, h path.Handlers) {
	pathGroup.POST("", h.Create())
	pathGroup.GET("", h.GetAll())
	pathGroup.GET("/:id", h.GetById())
	pathGroup.DELETE("/:id", h.Delete())
	pathGroup.PUT("/:id", h.Update())
}
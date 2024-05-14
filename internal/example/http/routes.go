package http

import (
	"github.com/esgi-challenge/backend/internal/example"
	"github.com/gin-gonic/gin"
)

func SetupExampleRoutes(exampleGroup *gin.RouterGroup, h example.Handlers) {
	exampleGroup.POST("", h.Create())
	exampleGroup.GET("", h.GetAll())
	exampleGroup.GET("/:id", h.GetById())
	exampleGroup.DELETE("/:id", h.Delete())
}

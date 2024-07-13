package http

import (
	"github.com/esgi-challenge/backend/internal/project"
	"github.com/gin-gonic/gin"
)

func SetupProjectRoutes(projectGroup *gin.RouterGroup, h project.Handlers) {
	projectGroup.POST("", h.Create())
	projectGroup.GET("", h.GetAll())
	projectGroup.GET("/:id", h.GetById())
	projectGroup.DELETE("/:id", h.Delete())
	projectGroup.PUT("/:id", h.Update())
	projectGroup.POST("/:id/join", h.JoinProject())
	projectGroup.DELETE("/:id/quit", h.QuitProject())
	projectGroup.GET("/:id/groups", h.GetGroups())
}

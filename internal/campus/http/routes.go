package http

import (
	"github.com/esgi-challenge/backend/internal/campus"
	"github.com/gin-gonic/gin"
)

func SetupCampusRoutes(campusGroup *gin.RouterGroup, h campus.Handlers) {
	campusGroup.POST("", h.Create())
	campusGroup.GET("", h.GetAll())
	campusGroup.GET("/:id", h.GetById())
	campusGroup.DELETE("/:id", h.Delete())
	campusGroup.PUT("/:id", h.Update())
}

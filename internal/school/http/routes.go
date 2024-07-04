package http

import (
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/gin-gonic/gin"
)

func SetupSchoolRoutes(schoolGroup *gin.RouterGroup, h school.Handlers) {
	schoolGroup.POST("", h.Create())
	schoolGroup.POST("/add/:kind", h.AddUser())
	schoolGroup.PUT("/update/:id", h.UpdateUser())
	schoolGroup.POST("/invite/:id", h.Invite())
	schoolGroup.GET("", h.GetByUser())
	schoolGroup.GET("/:id", h.GetById())
	schoolGroup.GET("/users/:kind", h.GetSchoolUsers())
	schoolGroup.DELETE("/remove/:kind/:id", h.RemoveUser())
	schoolGroup.DELETE("/:id", h.Delete())
}

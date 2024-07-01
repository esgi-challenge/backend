package http

import (
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/gin-gonic/gin"
)

func SetupSchoolRoutes(schoolGroup *gin.RouterGroup, h school.Handlers) {
	schoolGroup.POST("", h.Create())
	schoolGroup.POST("/:id", h.Invite())
	schoolGroup.GET("", h.GetByUser())
	schoolGroup.GET("/:id", h.GetById())
	schoolGroup.GET("/students", h.GetSchoolStudents())
	schoolGroup.DELETE("/:id", h.Delete())
}

package http

import (
	"github.com/esgi-challenge/backend/internal/school"
	"github.com/gin-gonic/gin"
)

func SetupSchoolRoutes(schoolGroup *gin.RouterGroup, h school.Handlers) {
	schoolGroup.POST("", h.Create())
	schoolGroup.POST("/student/add", h.AddUser())
	schoolGroup.PUT("/student/update/:id", h.UpdateUser())
	schoolGroup.POST("/invite/:id", h.Invite())
	schoolGroup.GET("", h.GetByUser())
	schoolGroup.GET("/:id", h.GetById())
	schoolGroup.GET("/students", h.GetSchoolStudents())
	schoolGroup.DELETE("/student/:id", h.RemoveStudent())
	schoolGroup.DELETE("/:id", h.Delete())
}

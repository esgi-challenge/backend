package http

import (
	"github.com/esgi-challenge/backend/internal/course"
	"github.com/gin-gonic/gin"
)

func SetupCourseRoutes(courseGroup *gin.RouterGroup, h course.Handlers) {
	courseGroup.POST("", h.Create())
	courseGroup.GET("", h.GetAll())
	courseGroup.GET("/:id", h.GetById())
	courseGroup.DELETE("/:id", h.Delete())
	courseGroup.PUT("/:id", h.Update())
}

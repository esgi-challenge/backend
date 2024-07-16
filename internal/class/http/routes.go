package http

import (
	"github.com/esgi-challenge/backend/internal/class"
	"github.com/gin-gonic/gin"
)

func SetupClassRoutes(classGroup *gin.RouterGroup, h class.Handlers) {
	classGroup.POST("/:id/add", h.Add())
	classGroup.DELETE("/:id/remove", h.Remove())
	classGroup.POST("", h.Create())
	classGroup.GET("", h.GetAll())
	classGroup.GET("/:id", h.GetById())
	classGroup.GET("/student", h.GetByStudent())
	classGroup.GET("/students/empty", h.GetClassLessStudents())
	classGroup.DELETE("/:id", h.Delete())
	classGroup.PUT("/:id", h.Update())
}

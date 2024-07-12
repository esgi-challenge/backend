package http

import (
	"github.com/esgi-challenge/backend/internal/schedule"
	"github.com/gin-gonic/gin"
)

func SetupScheduleRoutes(scheduleGroup *gin.RouterGroup, h schedule.Handlers) {
	scheduleGroup.POST("", h.Create())
	scheduleGroup.GET("", h.GetAll())
	scheduleGroup.GET("/unattended", h.GetUnattended())
	scheduleGroup.GET("/:id", h.GetById())
	scheduleGroup.GET("/:id/code", h.GetSignatureCode())
	scheduleGroup.DELETE("/:id", h.Delete())
	scheduleGroup.PUT("/:id", h.Update())
	scheduleGroup.POST("/:id/sign", h.Sign())
	scheduleGroup.GET("/:id/sign", h.CheckSign())
}

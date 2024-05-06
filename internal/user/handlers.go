package user

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create() gin.HandlerFunc
	// Update()
	// Delete()
	// GetByID()
}

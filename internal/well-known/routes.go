package wellknown

import (
	"github.com/gin-gonic/gin"
)

func SetupPathRoutes(pathGroup *gin.RouterGroup) {
	pathGroup.GET("/assetlinks.json", GetAssetLinks())
}

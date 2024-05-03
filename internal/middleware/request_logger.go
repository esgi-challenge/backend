package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (mw *MiddlewareManager) RequestMiddleware() gin.HandlerFunc {
  return func(ctx *gin.Context) {
    time := time.Now()

    req := ctx.Request

    mw.logger.Info("Server: Request received", 
      "method", req.Method,
      "url", req.URL,
      "host", req.Host,
      "remoteaddr", req.RemoteAddr,
      "time", time,
    ) 

    ctx.Next()
  }
}

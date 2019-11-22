package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jiaoww/gin-admin/internal/app/ginplus"
	"github.com/jiaoww/gin-admin/pkg/util"
)

// TraceMiddleware 跟踪ID中间件
func TraceMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		// 优先从请求头中获取请求ID，如果没有则使用UUID
		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = util.NewTraceID()
		}
		c.Set(ginplus.TraceIDKey, traceID)
		c.Next()
	}
}

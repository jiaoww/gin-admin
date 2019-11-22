package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/jiaoww/gin-admin/internal/app/config"
	"github.com/jiaoww/gin-admin/internal/app/errors"
	"github.com/jiaoww/gin-admin/internal/app/ginplus"
)

// CasbinMiddleware casbin中间件
func CasbinMiddleware(enforcer *casbin.SyncedEnforcer, skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := config.Global().Casbin
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		p := c.Request.URL.Path
		m := c.Request.Method
		if b, err := enforcer.Enforce(ginplus.GetUserID(c), p, m); err != nil {
			ginplus.ResError(c, errors.WithStack(err))
			return
		} else if !b {
			ginplus.ResError(c, errors.ErrNoPerm)
			return
		}
		c.Next()
	}
}

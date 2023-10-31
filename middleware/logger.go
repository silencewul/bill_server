package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InjectLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
	}
}

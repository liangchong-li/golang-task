package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware 日志
func LoggerMiddleware() gin.HandlerFunc {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		logger.WithFields(logrus.Fields{
			"status_code": params.StatusCode,
			"latency":     params.Latency,
			"client_ip":   params.ClientIP,
			"method":      params.Method,
			"path":        params.Path,
			"user_agent":  params.Request.UserAgent(),
			"error":       params.ErrorMessage,
			"timestamp":   params.TimeStamp.Format(time.RFC3339),
		}).Info("HTTP Request")

		return ""
	})
}

// ErrorHandleMiddleware 全局异常处理
func ErrorHandleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.WithFields(logrus.Fields{
					"error":  err,
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}).Error("Panic recovered")

				c.JSON(500, gin.H{
					"code":    500,
					"message": "Internal server error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

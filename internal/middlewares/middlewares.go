package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 日志中间件
// 日志格式: [GIN] 时间 | 客户端IP | 请求方法 | 状态码 | 延迟时间 | 请求路径
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()

		// 执行时间
		latency := end.Sub(start)

		// 请求信息
		method := c.Request.Method
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		path := c.Request.URL.Path

		// 打印日志
		// 格式: [GIN] 时间 | 客户端IP | 请求方法 | 状态码 | 延迟时间 | 请求路径
		fmt.Printf("[GIN] %s | %s | %s | %3d | %10v | %s\n",
			end.Format("2006-01-02 15:04:05"),
			clientIP,
			method,
			statusCode,
			latency,
			path,
		)
	}
}

// ErrorHandlerMiddleware 错误处理中间件
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查是否有错误
		if len(c.Errors) > 0 {
			c.JSON(-1, gin.H{
				"errors": c.Errors,
			})
			return
		}
	}
}

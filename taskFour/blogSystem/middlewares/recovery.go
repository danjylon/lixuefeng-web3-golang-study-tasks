package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
)

// RecoveryMiddleware 全局异常恢复中间件（相当于 @ControllerAdvice）
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 打印堆栈（用于调试）
				log.Printf("Panic: %v\n", err)
				for i := 2; ; i++ {
					_, file, line, ok := runtime.Caller(i)
					if !ok {
						break
					}
					log.Printf("  %s:%d", file, line)
				}
				// ✅ 安全的日志打印：区分不同类型
				var errMsg string
				switch v := err.(type) {
				case string:
					errMsg = v
				case error:
					errMsg = v.Error()
				default:
					errMsg = "unknown error"
				}
				//log.Printf("Recovery from panic: %s", errMsg)
				// 返回统一错误响应
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  errMsg,
				})
			}
		}()

		c.Next()
	}
}

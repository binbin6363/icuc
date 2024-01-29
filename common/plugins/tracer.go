package plugins

import (
	"math/rand"
	"time"

	"github.com/binbin6363/icuc/common/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// ZapTraceLogger 创建一个ZapTraceLogger中间件
func ZapTraceLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Gin Context 中获取 Trace ID（假设 Trace ID 存储在 Header 中）
		traceID := c.Request.Header.Get(log.LoggerTraceID)
		if len(traceID) == 0 {
			traceID = cast.ToString(rand.Uint64())
		}

		// 将 Trace ID 添加到 Zap Logger 的上下文字段中
		loggerWithTraceID := log.GetLogger().With(zap.String(log.LoggerTraceID, traceID))

		// 将 Zap Logger 添加到 Gin Context 中，以便在请求处理程序中使用
		c.Set(log.LoggerTag, loggerWithTraceID)

		// 继续处理请求
		t := time.Now()
		log.InfoContextf(c, "recv msg")
		c.Next()
		log.InfoContextf(c, "done, cost: %v", time.Since(t))
	}
}

package middleware

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"StakeBackend/internal/config"
	"StakeBackend/internal/redis"
)

// RateLimitMiddleware API限流中间件
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "rate_limit:" + ip
		max := config.GlobalConfig.RateLimit.MaxRequests
		window := config.GlobalConfig.RateLimit.WindowSeconds

		count, err := redis.RDB.Incr(redis.Ctx, key).Result()
		if err != nil {
			// Redis错误时降级处理：记录日志但允许请求通过
			c.Next()
			return
		}

		if count == 1 {
			redis.RDB.Expire(redis.Ctx, key, time.Duration(window)*time.Second)
		}

		if count > int64(max) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code": 429,
				"msg":  "请求过于频繁",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	// 访问跨域请求的http方法
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	// 允许跨域请求的头
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "Authorization"}
	// 允许跨域请求的源
	config.AllowOrigins = []string{"http://localhost:8080", "http://localhost:8000"}
	config.AllowCredentials = true
	return cors.New(config)
}

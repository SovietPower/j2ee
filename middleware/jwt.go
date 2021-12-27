package middleware

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/util"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = constant.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = constant.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != constant.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    constant.GetMsg(code),
				"data":   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

// JWTAdmin token验证中间件
func JWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = constant.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = constant.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = constant.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = constant.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else if claims.Authority != model.AuthorityAdmin {
				code = constant.ERROR_AUTH_INSUFFICIENT_AUTHORITY
			}
		}

		if code != constant.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    constant.GetMsg(code),
				"data":   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

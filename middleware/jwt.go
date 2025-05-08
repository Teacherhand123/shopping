package middleware

import (
	"shopping/pkg/e"
	"shopping/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")

		// 没有token
		if token == "" {
			code = 404
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				// 过期
				code = e.ErrorAuthCheckTokenTimeout
			}
		}

		if code != e.Success {
			// fmt.Println("来自middleware的json返回")
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

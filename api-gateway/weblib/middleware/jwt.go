package middleware

import (
	"api-gateway/pkg/util"
	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code uint32
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				code = 401
			}
		}
		if code != 200 {
			c.JSON(500, gin.H{
				"code": code,
				"msg":  "authorization failed",
			})

			c.Abort()
			return
		}
		c.Next()
	}
}

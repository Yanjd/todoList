package middleware

import (
	"api-gateway/pkg/logging"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(service []interface{}) gin.HandlerFunc {
	logging.Debug("init middleware")
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["userService"] = service[0]
		context.Next()
	}
}

func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logging.Debug("error middleware")
				context.JSON(200, gin.H{
					"code": 404,
					"msg":  r,
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}

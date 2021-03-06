package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		v1.POST("/user/register", handlers.UserRegister)
		v1.POST("/user/login", handlers.UserLogin)

		auth := v1.Group("/")
		auth.Use(middleware.Jwt())
		{
			auth.GET("tasks", handlers.GetTaskList)
			auth.POST("task", handlers.CreateTask)
			auth.GET("task/:id", handlers.GetTask)
			auth.PUT("task/:id", handlers.UpdateTask)
			auth.DELETE("task/:id", handlers.DeleteTask)
		}
	}

	return r
}

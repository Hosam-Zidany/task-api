package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/register", Register)
	r.POST("/login", Login)
	protected := r.Group("/api")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/me", Me)
		protected.POST("/tasks", CreateTask)
		protected.GET("/tasks", ListTasks)
		protected.GET("/tasks/:ID", GetTask)
		protected.PUT("/tasks/:ID", UpdateTask)
		protected.DELETE("/tasks/:ID", DeleteTask)
	}

	return r
}

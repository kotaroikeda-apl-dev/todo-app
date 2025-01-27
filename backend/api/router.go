package api

import (
	"todo/handlers"
	"todo/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// CORS ミドルウェアを適用
	r.Use(middlewares.CORSConfig())

	// API ルート
	r.GET("/api/tasks", handlers.GetTasks)
	r.POST("/api/tasks", handlers.CreateTask)
	r.PUT("/api/tasks/:id", handlers.UpdateTask)
	r.DELETE("/api/tasks/:id", handlers.DeleteTask)

	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})
}

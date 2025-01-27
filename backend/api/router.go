package api

import (
	"todo/controllers"
	"todo/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// CORS ミドルウェアを適用
	r.Use(middlewares.CORSConfig())

	// API ルート
	r.GET("/api/tasks", controllers.GetTasks)
	r.POST("/api/tasks", controllers.CreateTask)
	r.PUT("/api/tasks/:id", controllers.UpdateTask)
	r.DELETE("/api/tasks/:id", controllers.DeleteTask)

	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})
}

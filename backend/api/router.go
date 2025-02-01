package api

import (
	"todo/controllers"
	"todo/middlewares"
	"todo/repositories"
	"todo/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// CORS ミドルウェアを適用
	r.Use(middlewares.CORSConfig())

	// DIの実装
	repo := repositories.NewTaskRepository(db)
	service := services.NewTaskService(repo)
	controller := controllers.NewTaskController(service)

	// API ルート
	r.GET("/api/tasks", controller.GetTasks)
	r.POST("/api/tasks", controller.CreateTask)
	r.PUT("/api/tasks/:id", controller.UpdateTask)
	r.DELETE("/api/tasks/:id", controller.DeleteTask)

	// ヘルスチェック
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	return r
}

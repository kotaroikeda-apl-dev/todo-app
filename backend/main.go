package main

import (
	"todo/config"
	"todo/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://todo-elb-121795785.ap-northeast-1.elb.amazonaws.com",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	config.ConnectDB() // データベース接続

	// ルーティング設定
	r.GET("/api/tasks", handlers.GetTasks)
	r.POST("/api/tasks", handlers.CreateTask)
	r.PUT("/api/tasks/:id", handlers.UpdateTask)
	r.DELETE("/api/tasks/:id", handlers.DeleteTask)
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	r.Run(":8080")
}

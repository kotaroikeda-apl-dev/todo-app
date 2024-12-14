package main

import (
	"log"
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

	// データベース接続
	config.ConnectDB()

	// ルーティング設定
	r.GET("/api/tasks", handlers.GetTasks)
	r.POST("/api/tasks", handlers.CreateTask)
	r.PUT("/api/tasks/:id", handlers.UpdateTask)
	r.DELETE("/api/tasks/:id", handlers.DeleteTask)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	// HTTPS サーバー起動
	certFile := "/path/to/server.crt" // 証明書ファイルのパス
	keyFile := "/path/to/server.key"  // 秘密鍵ファイルのパス

	log.Println("Starting HTTPS server on port 8443...")
	err := r.RunTLS(":8443", certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to start HTTPS server: %v", err)
	}
}

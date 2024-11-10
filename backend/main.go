package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "todo/handlers"
    "todo/config"
)

func main() {
    r := gin.Default()

    // CORS設定
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
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

    r.Run(":8080")
}


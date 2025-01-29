package main

import (
	"todo/api"
	"todo/config"
	"todo/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// データベース接続
	config.ConnectDB()

	//corsチェック
	r.Use(middlewares.CORSConfig())
	// ルートの登録
	api.RegisterRoutes(r)

	r.Run(":8080")
}

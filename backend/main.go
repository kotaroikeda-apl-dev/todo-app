package main

import (
	"todo/api"
	"todo/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// データベース接続
	config.ConnectDB()

	// ルートの登録
	api.RegisterRoutes(r)

	r.Run(":8080")
}

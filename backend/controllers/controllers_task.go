package controllers

import (
	"net/http"
	"todo/config" // データベース設定をインポート
	"todo/models" // modelsパッケージをインポート

	"github.com/gin-gonic/gin"
)

// GETリクエストで全タスクを返すハンドラー
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	config.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// POSTリクエストで新しいタスクを追加するハンドラー
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}
func UpdateTask(c *gin.Context) {
	// 1. リクエストからデータを取得（バインド）
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. IDのパスパラメータを取得
	id := c.Param("id")
	var task models.Task

	// 3. IDでタスクを検索（存在確認）
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// 4. 取得したデータでフィールドを更新
	task.Title = input.Title
	task.Description = input.Description

	// 5. データベースに保存
	if err := config.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	// 6. 更新後のタスクをレスポンスとして返す
	c.JSON(http.StatusOK, task)
}
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	// IDでタスクを検索
	if err := config.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	config.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

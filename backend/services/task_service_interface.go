package services

import "todo/models"

type TaskService interface {
	GetAllTasks() ([]models.Task, error)
	CreateTask(task *models.Task) (*models.Task, error)
	UpdateTask(id int, title, description string) (*models.Task, error)
	DeleteTask(id int) error
}

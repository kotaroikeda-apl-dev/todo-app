package repositories

import "todo/models"

type TaskRepository interface {
	FindAllTasks() ([]models.Task, error)
	FindTaskByID(id int) (*models.Task, error)
	CreateTask(task *models.Task) (*models.Task, error)
	UpdateTask(task *models.Task) (*models.Task, error)
	DeleteTask(task *models.Task) error
}

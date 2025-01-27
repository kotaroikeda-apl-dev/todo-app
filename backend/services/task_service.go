package services

import (
	"todo/models"
	"todo/repositories"
)

func GetAllTasks() ([]models.Task, error) {
	return repositories.FindAllTasks()
}

func CreateTask(task *models.Task) (*models.Task, error) {
	return repositories.CreateTask(task)
}

func UpdateTask(id int, title, description string) (*models.Task, error) {
	task, err := repositories.FindTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.Title = title
	task.Description = description
	return repositories.UpdateTask(task)
}

func DeleteTask(id int) error {
	task, err := repositories.FindTaskByID(id)
	if err != nil {
		return err
	}
	return repositories.DeleteTask(task)
}

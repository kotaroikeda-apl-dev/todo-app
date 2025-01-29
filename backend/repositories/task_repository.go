package repositories

import (
	"todo/config"
	"todo/models"
)

func FindAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := config.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func CreateTask(task *models.Task) (*models.Task, error) {
	if err := config.DB.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func FindTaskByID(id int) (*models.Task, error) {
	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(task *models.Task) (*models.Task, error) {
	if err := config.DB.Save(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func DeleteTask(task *models.Task) error {
	return config.DB.Delete(task).Error
}

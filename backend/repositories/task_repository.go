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
	tx := config.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := config.DB.Create(task).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
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
	tx := config.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := config.DB.Save(task).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return task, nil
}

func DeleteTask(task *models.Task) error {
	tx := config.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := config.DB.Delete(task).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

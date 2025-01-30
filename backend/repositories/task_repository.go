package repositories

import (
	"time"
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
	task.UpdatedAt = time.Now().UTC().Truncate(time.Microsecond)

	sqlUpdateTask := `
		UPDATE tasks
		SET title = $1, description = $2, completed = $3, updated_at = $4
		WHERE id = $5
	`

	tx := config.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := tx.Exec(sqlUpdateTask, task.Title, task.Description, task.Completed, task.UpdatedAt, task.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return task, nil
}

func DeleteTask(task *models.Task) error {
	if err := config.DB.Delete(task).Error; err != nil {
		return err
	}
	return nil
}

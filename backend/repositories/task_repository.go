package repositories

import (
	"time"
	"todo/models"

	"gorm.io/gorm"
)

// taskRepository の実装
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository は TaskRepository のインスタンスを作成
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) FindAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *taskRepository) FindTaskByID(id int) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) UpdateTask(task *models.Task) (*models.Task, error) {
	task.UpdatedAt = time.Now().UTC().Truncate(time.Microsecond)

	sqlUpdateTask := `
		UPDATE tasks
		SET title = $1, description = $2, completed = $3, updated_at = $4
		WHERE id = $5
	`

	tx := r.db.Begin()
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

func (r *taskRepository) DeleteTask(task *models.Task) error {
	if err := r.db.Delete(task).Error; err != nil {
		return err
	}
	return nil
}

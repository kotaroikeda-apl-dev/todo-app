package services

import (
	"todo/models"
	"todo/repositories"
)

type taskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.FindAllTasks()
}

func (s *taskService) CreateTask(task *models.Task) (*models.Task, error) {
	return s.repo.CreateTask(task)
}

func (s *taskService) UpdateTask(id int, title, description string) (*models.Task, error) {
	task, err := s.repo.FindTaskByID(id)
	if err != nil {
		return nil, err
	}

	task.Title = title
	task.Description = description
	return s.repo.UpdateTask(task)
}

func (s *taskService) DeleteTask(id int) error {
	task, err := s.repo.FindTaskByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteTask(task)
}

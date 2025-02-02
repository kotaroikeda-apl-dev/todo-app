package services_test

import (
	"errors"
	"testing"
	"time"
	"todo/models"
	"todo/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// モックリポジトリの定義
type mockTaskRepository struct {
	mock.Mock
}

func (m *mockTaskRepository) FindAllTasks() ([]models.Task, error) {
	args := m.Called()
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *mockTaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *mockTaskRepository) FindTaskByID(id int) (*models.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *mockTaskRepository) UpdateTask(task *models.Task) (*models.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *mockTaskRepository) DeleteTask(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func TestGetAllTasks(t *testing.T) {
	mockRepo := new(mockTaskRepository)
	service := services.NewTaskService(mockRepo)

	tasks := []models.Task{
		{ID: 1, Title: "Task 1", Description: "Description 1"},
		{ID: 2, Title: "Task 2", Description: "Description 2"},
	}

	mockRepo.On("FindAllTasks").Return(tasks, nil)

	result, err := service.GetAllTasks()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Task 1", result[0].Title)
	assert.Equal(t, "Task 2", result[1].Title)
	assert.Equal(t, "Description 1", result[0].Description)
	assert.Equal(t, "Description 2", result[1].Description)
}

func TestCreateTask(t *testing.T) {
	mockRepo := new(mockTaskRepository)
	service := services.NewTaskService(mockRepo)

	task := &models.Task{
		Title:       "New Task",
		Description: "New Task Description",
		Completed:   false,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	mockRepo.On("CreateTask", task).Return(task, nil)

	result, err := service.CreateTask(task)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "New Task", result.Title)
	assert.Equal(t, "New Task Description", result.Description)
	assert.Equal(t, false, result.Completed)
}

func TestUpdateTask_Success(t *testing.T) {
	mockRepo := new(mockTaskRepository)
	service := services.NewTaskService(mockRepo)

	existingTask := &models.Task{
		ID:          1,
		Title:       "Old Title",
		Description: "Old Description",
	}

	updatedTask := &models.Task{
		ID:          1,
		Title:       "Updated Title",
		Description: "Updated Description",
	}

	mockRepo.On("FindTaskByID", 1).Return(existingTask, nil)
	mockRepo.On("UpdateTask", updatedTask).Return(updatedTask, nil)

	result, err := service.UpdateTask(1, "Updated Title", "Updated Description")
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", result.Title)
	assert.Equal(t, "Updated Description", result.Description)
}

func TestUpdateTask_NotFound(t *testing.T) {
	mockRepo := new(mockTaskRepository)
	service := services.NewTaskService(mockRepo)

	mockRepo.On("FindTaskByID", 1).Return((*models.Task)(nil), errors.New("not found"))

	result, err := service.UpdateTask(1, "Updated Title", "Updated Description")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestDeleteTask_Success(t *testing.T) {
	mockRepo := new(mockTaskRepository)
	service := services.NewTaskService(mockRepo)

	task := &models.Task{ID: 1}
	mockRepo.On("FindTaskByID", 1).Return(task, nil)
	mockRepo.On("DeleteTask", task).Return(nil)

	err := service.DeleteTask(1)
	assert.NoError(t, err)
}

func TestDeleteTask_NotFound(t *testing.T) {
	mockRepo := new(mockTaskRepository)
	service := services.NewTaskService(mockRepo)

	mockRepo.On("FindTaskByID", 1).Return((*models.Task)(nil), errors.New("not found"))

	err := service.DeleteTask(1)
	assert.Error(t, err)
}

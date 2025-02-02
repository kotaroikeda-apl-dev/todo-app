package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"todo/controllers"
	"todo/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// モックサービスの定義
type mockTaskService struct {
	mock.Mock
}

func (m *mockTaskService) GetAllTasks() ([]models.Task, error) {
	args := m.Called()
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *mockTaskService) CreateTask(task *models.Task) (*models.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *mockTaskService) UpdateTask(id int, title, description string) (*models.Task, error) {
	args := m.Called(id, title, description)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *mockTaskService) DeleteTask(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Gin のセットアップ
func setupTestContext() (*gin.Engine, *mockTaskService, *controllers.TaskController) {
	gin.SetMode(gin.TestMode)
	mockService := new(mockTaskService)
	controller := controllers.NewTaskController(mockService)

	r := gin.Default()
	r.GET("/tasks", controller.GetTasks)
	r.POST("/tasks", controller.CreateTask)
	r.PUT("/tasks/:id", controller.UpdateTask)
	r.DELETE("/tasks/:id", controller.DeleteTask)

	return r, mockService, controller
}

// GetTasks のテスト
func TestGetTasks_Success(t *testing.T) {
	r, mockService, _ := setupTestContext()

	tasks := []models.Task{
		{ID: 1, Title: "Task 1", Description: "Description 1"},
		{ID: 2, Title: "Task 2", Description: "Description 2"},
	}

	mockService.On("GetAllTasks").Return(tasks, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Task
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Len(t, response, 2)
	assert.Equal(t, "Task 1", response[0].Title)
	assert.Equal(t, "Description 1", response[0].Description)
	assert.Equal(t, "Task 2", response[1].Title)
	assert.Equal(t, "Description 2", response[1].Description)
}

func TestGetTasks_Failure(t *testing.T) {
	r, mockService, _ := setupTestContext()

	mockService.On("GetAllTasks").Return([]models.Task(nil), errors.New("failed to fetch tasks"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// CreateTask のテスト
func TestCreateTask_Success(t *testing.T) {
	r, mockService, _ := setupTestContext()

	task := &models.Task{
		ID:          1,
		Title:       "New Task",
		Description: "New Task Description",
	}

	mockService.On("CreateTask", mock.AnythingOfType("*models.Task")).Return(task, nil)

	taskJSON, _ := json.Marshal(task)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Task
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "New Task", response.Title)
	assert.Equal(t, "New Task Description", response.Description)
}

func TestCreateTask_Failure(t *testing.T) {
	r, mockService, _ := setupTestContext()

	mockService.On("CreateTask", mock.Anything).Return((*models.Task)(nil), errors.New("failed to create task"))

	taskJSON, _ := json.Marshal(models.Task{Title: "Invalid Task"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// UpdateTask のテスト
func TestUpdateTask_Success(t *testing.T) {
	r, mockService, _ := setupTestContext()

	taskID := 1
	updatedTask := &models.Task{
		ID:          uint(taskID),
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	mockService.On("UpdateTask", taskID, "Updated Task", "Updated Description").Return(updatedTask, nil)

	taskUpdateReq := models.TaskUpdateRequest{
		Title:       "Updated Task",
		Description: "Updated Description",
	}
	taskJSON, _ := json.Marshal(taskUpdateReq)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks/"+strconv.Itoa(taskID), bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Task
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Updated Task", response.Title)
	assert.Equal(t, "Updated Description", response.Description)
}

func TestUpdateTask_NotFound(t *testing.T) {
	r, mockService, _ := setupTestContext()

	taskID := 1
	mockService.On("UpdateTask", taskID, "Updated Task", "Updated Description").Return((*models.Task)(nil), errors.New("task not found"))

	taskUpdateReq := models.TaskUpdateRequest{
		Title:       "Updated Task",
		Description: "Updated Description",
	}
	taskJSON, _ := json.Marshal(taskUpdateReq)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks/"+strconv.Itoa(taskID), bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// DeleteTask のテスト
func TestDeleteTask_Success(t *testing.T) {
	r, mockService, _ := setupTestContext()

	taskID := 1
	mockService.On("DeleteTask", taskID).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(taskID), nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Task deleted", response["message"])
}

func TestDeleteTask_NotFound(t *testing.T) {
	r, mockService, _ := setupTestContext()

	taskID := 1
	mockService.On("DeleteTask", taskID).Return(errors.New("task not found"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(taskID), nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

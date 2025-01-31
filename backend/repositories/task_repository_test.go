package repositories_test

import (
	"testing"
	"time"
	"todo/config"
	"todo/models"
	"todo/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// テスト用にDBをモックする
func setupMockDB(t *testing.T) sqlmock.Sqlmock {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open mock GORM DB: %v", err)
	}

	config.DB = gormDB
	return mock
}

func TestFindAllTasks(t *testing.T) {
	mock := setupMockDB(t)

	mock.ExpectQuery(`SELECT \* FROM "tasks"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(1, "Task 1", "Description 1").
			AddRow(2, "Task 2", "Description 2"))

	tasks, err := repositories.FindAllTasks()
	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
}

func TestCreateTask(t *testing.T) {
	mock := setupMockDB(t)

	// `time.Now().UTC().Truncate(time.Microsecond)` で精度を統一
	now := time.Now().UTC().Truncate(time.Microsecond)

	task := &models.Task{
		Title:       "New Task",
		Description: "New Task Description",
		Completed:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "tasks" \("title","description","completed","created_at","updated_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5\) RETURNING "id"`).
		WithArgs(task.Title, task.Description, task.Completed, now, now). // 時刻を正確にマッチ
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))         // Query なので `WillReturnRows`

	mock.ExpectCommit()

	result, err := repositories.CreateTask(task)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "New Task", result.Title)
	assert.Equal(t, "New Task Description", result.Description)
	assert.Equal(t, false, result.Completed)
}

func TestFindTaskByID(t *testing.T) {
	mock := setupMockDB(t)

	mock.ExpectQuery(`SELECT \* FROM "tasks" WHERE "tasks"."id" = \$1 ORDER BY "tasks"."id" LIMIT \$2`).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow(1, "Task 1", "Description 1"))

	task, err := repositories.FindTaskByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, "Task 1", task.Title)
}

func TestUpdateTask(t *testing.T) {
	mock := setupMockDB(t)

	// `updated_at` の精度を統一
	now := time.Now().UTC().Truncate(time.Microsecond)

	task := &models.Task{
		ID:          1,
		Title:       "Updated Task",
		Description: "Updated Description",
		Completed:   false,
		UpdatedAt:   now,
	}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE tasks SET title = \$1, description = \$2, completed = \$3, updated_at = \$4 WHERE id = \$5`).
		WithArgs(task.Title, task.Description, task.Completed, sqlmock.AnyArg(), task.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	result, err := repositories.UpdateTask(task)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Updated Task", result.Title)
	assert.Equal(t, "Updated Description", result.Description)
	assert.Equal(t, false, result.Completed)
}

func TestDeleteTask(t *testing.T) {
	mock := setupMockDB(t)

	task := &models.Task{ID: 1}
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "tasks" WHERE "tasks"."id" = \$1`).WithArgs(task.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repositories.DeleteTask(task)
	assert.NoError(t, err)
}

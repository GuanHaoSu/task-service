package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task_service/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() (*gin.Engine, *TaskHandler) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	handler := NewTaskHandler()
	return r, handler
}

func TestGetTasks(t *testing.T) {
	r, handler := setupTestRouter()
	r.GET("/tasks", handler.GetTasks)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateTask(t *testing.T) {
	t.Run("valid task", func(t *testing.T) {
		r, handler := setupTestRouter()
		r.POST("/tasks", handler.CreateTask)

		task := models.Task{
			Name:   "Test Task",
			Status: 0,
		}
		jsonValue, _ := json.Marshal(task)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response models.Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.ID)
		assert.Equal(t, task.Name, response.Name)
		assert.Equal(t, task.Status, response.Status)
	})

	t.Run("invalid status", func(t *testing.T) {
		r, handler := setupTestRouter()
		r.POST("/tasks", handler.CreateTask)

		invalidTask := struct {
			Name   string `json:"name"`
			Status int    `json:"status"`
		}{
			Name:   "Test Task",
			Status: 2,
		}
		jsonValue, _ := json.Marshal(invalidTask)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("invalid json", func(t *testing.T) {
		r, handler := setupTestRouter()
		r.POST("/tasks", handler.CreateTask)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte(`{invalid json}`)))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("valid update", func(t *testing.T) {
		r, handler := setupTestRouter()
		r.PUT("/tasks/:id", handler.UpdateTask)

		task := &models.Task{
			ID:     "test-id",
			Name:   "Test Task",
			Status: 0,
		}
		handler.store.Create(task)

		updatedTask := models.Task{
			Name:   "Updated Task",
			Status: 1,
		}
		jsonValue, _ := json.Marshal(updatedTask)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/tasks/test-id", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Task
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "test-id", response.ID)
		assert.Equal(t, updatedTask.Name, response.Name)
		assert.Equal(t, updatedTask.Status, response.Status)
	})

	t.Run("invalid status", func(t *testing.T) {
		r, handler := setupTestRouter()
		r.PUT("/tasks/:id", handler.UpdateTask)

		task := &models.Task{
			ID:     "test-id",
			Name:   "Test Task",
			Status: 0,
		}
		handler.store.Create(task)

		invalidTask := struct {
			Name   string `json:"name"`
			Status int    `json:"status"`
		}{
			Name:   "Updated Task",
			Status: 2, // Invalid status
		}
		jsonValue, _ := json.Marshal(invalidTask)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/tasks/test-id", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("invalid json", func(t *testing.T) {
		r, handler := setupTestRouter()
		r.PUT("/tasks/:id", handler.UpdateTask)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/tasks/test-id", bytes.NewBuffer([]byte(`{invalid json}`)))
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDeleteTask(t *testing.T) {
	r, handler := setupTestRouter()
	r.DELETE("/tasks/:id", handler.DeleteTask)

	// First create a task
	task := &models.Task{
		ID:     "test-id",
		Name:   "Test Task",
		Status: 0,
	}
	handler.store.Create(task)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/test-id", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

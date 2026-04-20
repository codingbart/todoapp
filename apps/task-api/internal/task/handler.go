package task

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/codingbart/todoapp/task-api/internal/response"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// @Router /users/{userId}/tasks [get]
// @Tags tasks
// @Summary List all tasks for a user
// @Security OAuth2
// @Param userId path string true "User ID"
// @Param status query string false "Filter by status (todo, in_progress, done)"
// @Param priority query string false "Filter by priority (low, medium, high)"
// @Success 200 {object} response.Response[[]TaskResponse]
// @Failure 400 {object} response.ErrorResponse
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(chi.URLParam(r, "userId"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	filters := TaskFilters{
		Status:   r.URL.Query().Get("status"),
		Priority: r.URL.Query().Get("priority"),
	}

	tasks, err := h.service.GetAll(r.Context(), userID, filters)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch tasks")
		return
	}

	response.Write(w, http.StatusOK, tasks)
}

// @Router /users/{userId}/dashboard [get]
// @Tags dashboard
// @Summary Get task summary dashboard for a user
// @Security OAuth2
// @Param userId path string true "User ID"
// @Success 200 {object} response.Response[DashboardResponse]
// @Failure 400 {object} response.ErrorResponse
func (h *Handler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(chi.URLParam(r, "userId"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	dashboard, err := h.service.GetDashboard(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch dashboard")
		return
	}

	response.Write(w, http.StatusOK, dashboard)
}

// @Router /users/{userId}/tasks/{id} [get]
// @Tags tasks
// @Summary Get a task by ID
// @Security OAuth2
// @Param userId path string true "User ID"
// @Param id path string true "Task ID"
// @Success 200 {object} response.Response[TaskResponse]
// @Failure 404 {object} response.ErrorResponse
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	task, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(w, http.StatusNotFound, "Task not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, "Failed to fetch task")
		return
	}

	response.Write(w, http.StatusOK, task)
}

// @Router /users/{userId}/tasks [post]
// @Tags tasks
// @Summary Create a new task
// @Security OAuth2
// @Param userId path string true "User ID"
// @Param body body CreateTaskRequest true "Task data"
// @Success 201 {object} response.Response[TaskResponse]
// @Failure 400 {object} response.ErrorResponse
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(chi.URLParam(r, "userId"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Title == "" {
		response.Error(w, http.StatusBadRequest, "Title is required")
		return
	}

	task, err := h.service.Create(r.Context(), userID, req)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to create task")
		return
	}

	response.Write(w, http.StatusCreated, task)
}

// @Router /users/{userId}/tasks/{id} [put]
// @Tags tasks
// @Summary Update a task
// @Security OAuth2
// @Param userId path string true "User ID"
// @Param id path string true "Task ID"
// @Param body body UpdateTaskRequest true "Task data"
// @Success 200 {object} response.Response[TaskResponse]
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Title == "" {
		response.Error(w, http.StatusBadRequest, "Title is required")
		return
	}

	task, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(w, http.StatusNotFound, "Task not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, "Failed to update task")
		return
	}

	response.Write(w, http.StatusOK, task)
}

// @Router /users/{userId}/tasks/{id} [delete]
// @Tags tasks
// @Summary Delete a task
// @Security OAuth2
// @Param userId path string true "User ID"
// @Param id path string true "Task ID"
// @Success 204
// @Failure 400 {object} response.ErrorResponse
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid task ID")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete task")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

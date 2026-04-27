package task

import (
	"context"
	"database/sql"
	"time"

	db "github.com/codingbart/todoapp/task-api/internal/db/postgresql"
	"github.com/google/uuid"
)

type Service interface {
	GetAll(ctx context.Context, userID uuid.UUID, filters TaskFilters) ([]TaskResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (TaskResponse, error)
	Create(ctx context.Context, userID uuid.UUID, req CreateTaskRequest) (TaskResponse, error)
	Update(ctx context.Context, id uuid.UUID, req UpdateTaskRequest) (TaskResponse, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetDashboard(ctx context.Context, userID uuid.UUID) (DashboardResponse, error)
}

type service struct {
	queries *db.Queries
}

func NewService(queries *db.Queries) Service {
	return &service{queries: queries}
}

func (s *service) GetAll(ctx context.Context, userID uuid.UUID, filters TaskFilters) ([]TaskResponse, error) {
	var tasks []db.Task
	var err error

	switch {
	case filters.Status != "" && filters.Priority != "":
		// Both filters — get by status then filter by priority in-memory
		tasks, err = s.queries.FindAllTasksByUserIdAndStatus(ctx, db.FindAllTasksByUserIdAndStatusParams{
			UserID: userID,
			Status: db.TaskStatus(filters.Status),
		})
		if err == nil {
			filtered := tasks[:0]
			for _, t := range tasks {
				if t.Priority == db.TaskPriority(filters.Priority) {
					filtered = append(filtered, t)
				}
			}
			tasks = filtered
		}
	case filters.Status != "":
		tasks, err = s.queries.FindAllTasksByUserIdAndStatus(ctx, db.FindAllTasksByUserIdAndStatusParams{
			UserID: userID,
			Status: db.TaskStatus(filters.Status),
		})
	case filters.Priority != "":
		tasks, err = s.queries.FindAllTasksByUserIdAndPriority(ctx, db.FindAllTasksByUserIdAndPriorityParams{
			UserID:   userID,
			Priority: db.TaskPriority(filters.Priority),
		})
	default:
		tasks, err = s.queries.FindAllTasksByUserId(ctx, userID)
	}

	if err != nil {
		return nil, err
	}

	result := make([]TaskResponse, 0, len(tasks))
	for _, t := range tasks {
		result = append(result, toTaskResponse(t))
	}
	return result, nil
}

func (s *service) GetDashboard(ctx context.Context, userID uuid.UUID) (DashboardResponse, error) {
	rows, err := s.queries.CountTasksByUserIdGroupedByStatus(ctx, userID)
	if err != nil {
		return DashboardResponse{}, err
	}

	var dash DashboardResponse
	for _, row := range rows {
		switch row.Status {
		case db.TaskStatusTodo:
			dash.Todo = row.Count
		case db.TaskStatusInProgress:
			dash.InProgress = row.Count
		case db.TaskStatusDone:
			dash.Done = row.Count
		}
		dash.Total += row.Count
	}
	return dash, nil
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (TaskResponse, error) {
	t, err := s.queries.FindTaskById(ctx, id)
	if err != nil {
		return TaskResponse{}, err
	}
	return toTaskResponse(t), nil
}

func (s *service) Create(ctx context.Context, userID uuid.UUID, req CreateTaskRequest) (TaskResponse, error) {
	params := db.CreateTaskParams{
		UserID:   userID,
		Title:    req.Title,
		Status:   db.TaskStatus(req.Status),
		Priority: db.TaskPriority(req.Priority),
	}

	if req.Description != nil {
		params.Description = sql.NullString{String: *req.Description, Valid: true}
	}

	if req.DueDate != nil {
		t, err := time.Parse(time.DateOnly, *req.DueDate)
		if err != nil {
			return TaskResponse{}, err
		}
		params.DueDate = sql.NullTime{Time: t, Valid: true}
	}

	task, err := s.queries.CreateTask(ctx, params)
	if err != nil {
		return TaskResponse{}, err
	}
	return toTaskResponse(task), nil
}

func (s *service) Update(ctx context.Context, id uuid.UUID, req UpdateTaskRequest) (TaskResponse, error) {
	params := db.UpdateTaskParams{
		ID:       id,
		Title:    req.Title,
		Status:   db.TaskStatus(req.Status),
		Priority: db.TaskPriority(req.Priority),
	}

	if req.Description != nil {
		params.Description = sql.NullString{String: *req.Description, Valid: true}
	}

	if req.DueDate != nil {
		t, err := time.Parse(time.DateOnly, *req.DueDate)
		if err != nil {
			return TaskResponse{}, err
		}
		params.DueDate = sql.NullTime{Time: t, Valid: true}
	}

	task, err := s.queries.UpdateTask(ctx, params)
	if err != nil {
		return TaskResponse{}, err
	}
	return toTaskResponse(task), nil
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	rows, err := s.queries.DeleteTaskById(ctx, id)
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func toTaskResponse(t db.Task) TaskResponse {
	resp := TaskResponse{
		ID:        t.ID,
		Title:     t.Title,
		Status:    string(t.Status),
		Priority:  string(t.Priority),
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}

	if t.Description.Valid {
		resp.Description = &t.Description.String
	}

	if t.DueDate.Valid {
		resp.DueDate = &t.DueDate.Time
	}

	return resp
}

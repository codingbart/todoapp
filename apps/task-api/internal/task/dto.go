package task

import (
	"errors"
	"time"

	"github.com/codingbart/todoapp/task-api/internal/response"
	"github.com/google/uuid"
)

type TaskListResponse = response.Response[[]TaskResponse]
type TaskSingleResponse = response.Response[TaskResponse]
type DashboardSingleResponse = response.Response[DashboardResponse]

type CreateTaskRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Status      string  `json:"status"`
	Priority    string  `json:"priority"`
	DueDate     *string `json:"due_date,omitempty"`
}

type UpdateTaskRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Status      string  `json:"status"`
	Priority    string  `json:"priority"`
	DueDate     *string `json:"due_date,omitempty"`
}

type TaskResponse struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description,omitempty"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskFilters struct {
	Status   string
	Priority string
}

type DashboardResponse struct {
	Todo       int64 `json:"todo"`
	InProgress int64 `json:"in_progress"`
	Done       int64 `json:"done"`
	Total      int64 `json:"total"`
}

var (
	validStatuses   = map[string]struct{}{"todo": {}, "in_progress": {}, "done": {}}
	validPriorities = map[string]struct{}{"low": {}, "medium": {}, "high": {}}
)

func validateStatus(s string) error {
	if _, ok := validStatuses[s]; !ok {
		return errors.New("status must be one of: todo, in_progress, done")
	}
	return nil
}

func validatePriority(p string) error {
	if _, ok := validPriorities[p]; !ok {
		return errors.New("priority must be one of: low, medium, high")
	}
	return nil
}

func validateDueDate(d *string) error {
	if d == nil || *d == "" {
		return nil
	}
	if _, err := time.Parse(time.DateOnly, *d); err != nil {
		return errors.New("due_date must be in YYYY-MM-DD format")
	}
	return nil
}

func (r CreateTaskRequest) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if err := validateStatus(r.Status); err != nil {
		return err
	}
	if err := validatePriority(r.Priority); err != nil {
		return err
	}
	return validateDueDate(r.DueDate)
}

func (r UpdateTaskRequest) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if err := validateStatus(r.Status); err != nil {
		return err
	}
	if err := validatePriority(r.Priority); err != nil {
		return err
	}
	return validateDueDate(r.DueDate)
}

func (f TaskFilters) Validate() error {
	if f.Status != "" {
		if err := validateStatus(f.Status); err != nil {
			return err
		}
	}
	if f.Priority != "" {
		if err := validatePriority(f.Priority); err != nil {
			return err
		}
	}
	return nil
}

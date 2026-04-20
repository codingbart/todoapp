package task_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/codingbart/todoapp/task-api/internal/task"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type mockService struct {
	tasks     []task.TaskResponse
	task      task.TaskResponse
	dashboard task.DashboardResponse
	err       error

	capturedFilters task.TaskFilters
}

func (m *mockService) GetAll(_ context.Context, _ uuid.UUID, filters task.TaskFilters) ([]task.TaskResponse, error) {
	m.capturedFilters = filters
	return m.tasks, m.err
}

func (m *mockService) GetByID(_ context.Context, _ uuid.UUID) (task.TaskResponse, error) {
	return m.task, m.err
}

func (m *mockService) Create(_ context.Context, _ uuid.UUID, _ task.CreateTaskRequest) (task.TaskResponse, error) {
	return m.task, m.err
}

func (m *mockService) Update(_ context.Context, _ uuid.UUID, _ task.UpdateTaskRequest) (task.TaskResponse, error) {
	return m.task, m.err
}

func (m *mockService) Delete(_ context.Context, _ uuid.UUID) error {
	return m.err
}

func (m *mockService) GetDashboard(_ context.Context, _ uuid.UUID) (task.DashboardResponse, error) {
	return m.dashboard, m.err
}

func withChiParams(r *http.Request, params map[string]string) *http.Request {
	rctx := chi.NewRouteContext()
	for k, v := range params {
		rctx.URLParams.Add(k, v)
	}
	ctx := context.WithValue(r.Context(), chi.RouteCtxKey, rctx)
	return r.WithContext(ctx)
}

// --- GetAll ---

func TestGetAll_Success(t *testing.T) {
	userID := uuid.New()
	tasks := []task.TaskResponse{
		{ID: uuid.New(), Title: "Test task", Status: "todo", Priority: "low", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	svc := &mockService{tasks: tasks}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userID.String()+"/tasks", nil)
	req = withChiParams(req, map[string]string{"userId": userID.String()})
	rec := httptest.NewRecorder()

	h.GetAll(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}

func TestGetAll_InvalidUserID(t *testing.T) {
	svc := &mockService{}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/users/invalid/tasks", nil)
	req = withChiParams(req, map[string]string{"userId": "invalid"})
	rec := httptest.NewRecorder()

	h.GetAll(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

func TestGetAll_FilterByStatus(t *testing.T) {
	userID := uuid.New()
	svc := &mockService{tasks: []task.TaskResponse{}}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userID.String()+"/tasks?status=done", nil)
	req = withChiParams(req, map[string]string{"userId": userID.String()})
	rec := httptest.NewRecorder()

	h.GetAll(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if svc.capturedFilters.Status != "done" {
		t.Fatalf("expected status filter 'done', got %q", svc.capturedFilters.Status)
	}
}

func TestGetAll_FilterByPriority(t *testing.T) {
	userID := uuid.New()
	svc := &mockService{tasks: []task.TaskResponse{}}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userID.String()+"/tasks?priority=high", nil)
	req = withChiParams(req, map[string]string{"userId": userID.String()})
	rec := httptest.NewRecorder()

	h.GetAll(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if svc.capturedFilters.Priority != "high" {
		t.Fatalf("expected priority filter 'high', got %q", svc.capturedFilters.Priority)
	}
}

func TestGetAll_FilterByStatusAndPriority(t *testing.T) {
	userID := uuid.New()
	svc := &mockService{tasks: []task.TaskResponse{}}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userID.String()+"/tasks?status=todo&priority=low", nil)
	req = withChiParams(req, map[string]string{"userId": userID.String()})
	rec := httptest.NewRecorder()

	h.GetAll(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if svc.capturedFilters.Status != "todo" || svc.capturedFilters.Priority != "low" {
		t.Fatalf("expected filters status=todo priority=low, got status=%q priority=%q", svc.capturedFilters.Status, svc.capturedFilters.Priority)
	}
}

// --- GetByID ---

func TestGetByID_Success(t *testing.T) {
	userID := uuid.New()
	taskID := uuid.New()
	resp := task.TaskResponse{ID: taskID, Title: "Test", Status: "todo", Priority: "low", CreatedAt: time.Now(), UpdatedAt: time.Now()}

	svc := &mockService{task: resp}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userID.String()+"/tasks/"+taskID.String(), nil)
	req = withChiParams(req, map[string]string{"userId": userID.String(), "id": taskID.String()})
	rec := httptest.NewRecorder()

	h.GetByID(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}

func TestGetByID_InvalidID(t *testing.T) {
	userID := uuid.New()
	svc := &mockService{}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userID.String()+"/tasks/invalid", nil)
	req = withChiParams(req, map[string]string{"userId": userID.String(), "id": "invalid"})
	rec := httptest.NewRecorder()

	h.GetByID(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

// --- Create ---

func TestCreate_Success(t *testing.T) {
	userID := uuid.New()
	taskResp := task.TaskResponse{ID: uuid.New(), Title: "New task", Status: "todo", Priority: "medium", CreatedAt: time.Now(), UpdatedAt: time.Now()}

	svc := &mockService{task: taskResp}
	h := task.NewHandler(svc)

	body, _ := json.Marshal(task.CreateTaskRequest{Title: "New task", Status: "todo", Priority: "medium"})
	req := httptest.NewRequest(http.MethodPost, "/users/"+userID.String()+"/tasks", bytes.NewReader(body))
	req = withChiParams(req, map[string]string{"userId": userID.String()})
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}
}

func TestCreate_InvalidUserID(t *testing.T) {
	svc := &mockService{}
	h := task.NewHandler(svc)

	body, _ := json.Marshal(task.CreateTaskRequest{Title: "New task", Status: "todo", Priority: "medium"})
	req := httptest.NewRequest(http.MethodPost, "/users/invalid/tasks", bytes.NewReader(body))
	req = withChiParams(req, map[string]string{"userId": "invalid"})
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

func TestCreate_MissingTitle(t *testing.T) {
	userID := uuid.New()
	svc := &mockService{}
	h := task.NewHandler(svc)

	body, _ := json.Marshal(task.CreateTaskRequest{Status: "todo", Priority: "medium"})
	req := httptest.NewRequest(http.MethodPost, "/users/"+userID.String()+"/tasks", bytes.NewReader(body))
	req = withChiParams(req, map[string]string{"userId": userID.String()})
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

func TestCreate_InvalidBody(t *testing.T) {
	userID := uuid.New()
	svc := &mockService{}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodPost, "/users/"+userID.String()+"/tasks", bytes.NewReader([]byte("invalid")))
	req = withChiParams(req, map[string]string{"userId": userID.String()})
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

// --- Update ---

func TestUpdate_Success(t *testing.T) {
	userID := uuid.New()
	taskID := uuid.New()
	taskResp := task.TaskResponse{ID: taskID, Title: "Updated", Status: "done", Priority: "high", CreatedAt: time.Now(), UpdatedAt: time.Now()}

	svc := &mockService{task: taskResp}
	h := task.NewHandler(svc)

	body, _ := json.Marshal(task.UpdateTaskRequest{Title: "Updated", Status: "done", Priority: "high"})
	req := httptest.NewRequest(http.MethodPut, "/users/"+userID.String()+"/tasks/"+taskID.String(), bytes.NewReader(body))
	req = withChiParams(req, map[string]string{"userId": userID.String(), "id": taskID.String()})
	rec := httptest.NewRecorder()

	h.Update(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}

func TestUpdate_InvalidID(t *testing.T) {
	userID := uuid.New()
	svc := &mockService{}
	h := task.NewHandler(svc)

	body, _ := json.Marshal(task.UpdateTaskRequest{Title: "Updated", Status: "done", Priority: "high"})
	req := httptest.NewRequest(http.MethodPut, "/users/"+userID.String()+"/tasks/invalid", bytes.NewReader(body))
	req = withChiParams(req, map[string]string{"userId": userID.String(), "id": "invalid"})
	rec := httptest.NewRecorder()

	h.Update(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

// --- Delete ---

func TestDelete_Success(t *testing.T) {
	userID := uuid.New()
	taskID := uuid.New()
	svc := &mockService{}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodDelete, "/users/"+userID.String()+"/tasks/"+taskID.String(), nil)
	req = withChiParams(req, map[string]string{"userId": userID.String(), "id": taskID.String()})
	rec := httptest.NewRecorder()

	h.Delete(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", rec.Code)
	}
}

func TestDelete_InvalidID(t *testing.T) {
	userID := uuid.New()
	svc := &mockService{}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodDelete, "/users/"+userID.String()+"/tasks/invalid", nil)
	req = withChiParams(req, map[string]string{"userId": userID.String(), "id": "invalid"})
	rec := httptest.NewRecorder()

	h.Delete(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

// --- Dashboard ---

func TestGetDashboard_Success(t *testing.T) {
	userID := uuid.New()
	svc := &mockService{
		dashboard: task.DashboardResponse{Todo: 3, InProgress: 2, Done: 5, Total: 10},
	}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/users/"+userID.String()+"/dashboard", nil)
	req = withChiParams(req, map[string]string{"userId": userID.String()})
	rec := httptest.NewRecorder()

	h.GetDashboard(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resp struct {
		Data task.DashboardResponse `json:"data"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp.Data.Total != 10 {
		t.Fatalf("expected total 10, got %d", resp.Data.Total)
	}
}

func TestGetDashboard_InvalidUserID(t *testing.T) {
	svc := &mockService{}
	h := task.NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/users/invalid/dashboard", nil)
	req = withChiParams(req, map[string]string{"userId": "invalid"})
	rec := httptest.NewRecorder()

	h.GetDashboard(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rec.Code)
	}
}

package health

import (
	"net/http"

	"github.com/codingbart/todoapp/task-api/internal/response"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	status := h.service.GetHealthStatus()
	response.Write(w, http.StatusOK, status)
}

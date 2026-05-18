package health

type HealthResponse struct {
	Status string `json:"status" validate:"required"`
}

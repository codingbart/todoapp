package health

type Service interface {
	GetHealthStatus() HealthResponse
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) GetHealthStatus() HealthResponse {
	return HealthResponse{Status: "ok"}
}

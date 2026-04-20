package response

type ErrorResponse struct {
	Error string `json:"error"`
}

type Response[T any] struct {
	Data T `json:"data"`
}

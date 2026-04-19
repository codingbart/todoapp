package response

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, body any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(body)
}

func Write[T any](w http.ResponseWriter, status int, data T) error {
	return writeJSON(w, status, Response[T]{
		Data: data,
	})
}

func Error(w http.ResponseWriter, status int, code, message string) error {
	return writeJSON(w, status, ErrorResponse{
		Error: ErrorBody{
			Code:    code,
			Message: message,
		},
	})
}

package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Success bool `json:"success"`
	Data any `json:"data"`
}

type ErrorResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

func WriteJSON(
	w http.ResponseWriter,
	status int,
	data any,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
func Success(
	w http.ResponseWriter,
	status int,
	data any,
) {

	resp := SuccessResponse{
		Success: true,
		Data: data,
	}

	WriteJSON(w, status, resp)

}

func Error(
	w http.ResponseWriter,
	status int,
	message string,
) {

	resp := ErrorResponse{
		Success: false,
		Message: message,
	}

	WriteJSON(w, status, resp)

}

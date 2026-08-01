package response

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type SuccessResponse struct {
	Success bool `json:"success"`
	Data any `json:"data"`
	Errors  map[string]string `json:"errors,omitempty"`
}

type ErrorResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func writeJSON(
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
		Errors: nil,
	}

	writeJSON(w, status, resp)

}

func Error(
	w http.ResponseWriter,
	status int,
	message string,
) {

	resp := ErrorResponse{
		Success: false,
		Message: message,
		Errors:  nil,
	}

	writeJSON(w, status, resp)

}

func ValidationError(
	w http.ResponseWriter,
	err error,
) {
	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {

		switch e.Tag() {

		case "required":
			errors[e.Field()] = "This field is required"

		case "email":
			errors[e.Field()] = "Invalid email address"

		case "min":
			errors[e.Field()] = "Minimum length is " + e.Param()

		case "max":
			errors[e.Field()] = "Maximum length is " + e.Param()

		default:
			errors[e.Field()] = "Invalid value"
		}
	}

	resp := ErrorResponse{
		Success: false,
		Message: "Validation failed",
		Errors:  errors,
	}

	writeJSON(w, http.StatusBadRequest, resp)
}

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/dto"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/response"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/service"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/validator"
)
type AuthHandler struct{
	authService *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService:service,
	}
}

// Register handles user registration requests
func (authHandler *AuthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", authHandler.Register)

}
func (authHandler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerDTO dto.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&registerDTO); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if err := validator.Validate.Struct(registerDTO); err != nil {
		response.ValidationError(w, err)
		return
	}

	user, err := authHandler.authService.Register(r.Context(), registerDTO.ToModel())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, user)
}
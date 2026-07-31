package routes

import (
	"net/http"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/handler"
)

func SetupRoutes(userHandler *handler.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handler.HomeHandler)
	mux.HandleFunc("GET /health", handler.Health)

	//users
	mux.HandleFunc("GET /users", userHandler.GetUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUser)
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)

	return mux
}

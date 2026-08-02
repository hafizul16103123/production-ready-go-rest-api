package routes

import (
	"net/http"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/handler"
)

func SetupRoutes(userHandler *handler.UserHandler, authHandler *handler.AuthHandler) *http.ServeMux {
	mux := http.NewServeMux()
	// mux.HandleFunc("GET /", handler.HomeHandler)
	// mux.HandleFunc("GET /health", handler.Health)

	//users
	userHandler.RegisterRoutes(mux)
	authHandler.RegisterRoutes(mux)

	return mux
}

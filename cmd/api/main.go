package main

import (
	"context"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/config"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/handler"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/repository"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/service"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/middleware"
	"github.com/hafizul16103123/production-ready-go-rest-api/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("invalid config: %v", err)
	}

	repo := repository.NewMemoryRepository()
	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)

	mux := routes.SetupRoutes(userHandler)

	handler := middleware.Chain(
		mux,
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
		middleware.RecoveryMiddleware,
		middleware.RequestIdMiddleware,
	)

	// Configure production-ready HTTP server
	server := &http.Server{
		Addr:           ":" + cfg.Port,
		Handler:        handler,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20, // Limit request header size (1 MB)
	}

	// Start the HTTP server in a separate goroutine
	// so the main goroutine can listen for shutdown signals.
	go func() {
		// Start accepting incoming HTTP requests.
		// ListenAndServe() blocks until the server stops.
		if err := server.ListenAndServe(); err != nil {
			// Ignore the expected error returned after Shutdown().
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	// Print startup message
	log.Println("Server started on http://localhost:" + cfg.Port)

	// Create a channel to receive OS signals
	// Notify this channel when Ctrl+C or SIGTERM is received
	// Block until a shutdown signal arrives

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")

	// Give existing requests up to 5 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	// Stops accepting new requests and waits for active ones.
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server gracefully shutdown")
}

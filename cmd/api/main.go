package main

/*

| Signal  | Meaning                             |
| ------- | ----------------------------------- |
| SIGINT  | Ctrl+C                              |
| SIGTERM | Stop application gracefully         |
| SIGKILL | Kill immediately (cannot be caught) |

*/

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/config"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/database"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/handler"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/logger"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/middleware"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/repository"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/service"
	"github.com/hafizul16103123/production-ready-go-rest-api/routes"
)

func main() {
	// Config
	cfg, err := config.MustLoad()
	if err != nil {
		slog.Error("invalid config", "error", err)
		os.Exit(1)
	}
	
	//Logger
	logger.Init()

	// Database
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(1)
	}
	slog.Info("Database connected successfully")
	defer db.Close()


	// repo := repository.NewMemoryRepository()
	userRepo := repository.NewPostgresRepository(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

	mux := routes.SetupRoutes(userHandler, authHandler)

	/*

		আপনার Chain() Function-এর Golden Rule:
		Chain()-এ Middleware যেই Order-এ লিখবেন, Request সেই একই Order-এ Execute হবে।
		কারণ Chain() Reverse Loop ব্যবহার করে Chain Build করে, কিন্তু Execution Order ঠিক developer যে Order দিয়েছেন সেটাই বজায় রাখে।

		প্রতিটি Middleware তার পরের Handler-কে Wrap (ঘিরে) করে। তাই Chain তৈরি করার সময় শেষ Middleware থেকে প্রথম Middleware পর্যন্ত (last → first) তৈরি করা হয়।
		এর ফলে Chain()-এ যে Middleware-টি প্রথমে দেওয়া হয়, সেটিই সবচেয়ে বাইরের Wrapper হয় এবং Request আসলে সেটিই সবার আগে Execute হয়।

		Order Should be:
		Recovery
		Request ID
		Logging
		Timeout
		Rate Limiting
		CORS
		Authentication
		Authorization
		Validation
		Business Handler

	*/

	handler := middleware.Chain(
		mux,
		middleware.RecoveryMiddleware, // সবচেয়ে বাইরের Wrapper হয় এবং Request আসলে সেটই সবার আগে Execute হয়।
		middleware.RequestIdMiddleware,
		middleware.LoggingMiddleware,
		middleware.AuthMiddleware,
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
				slog.Error("Server error", "error", err)
				os.Exit(1)
			}
		}
	}()

	// Print startup message
	slog.Info("Server started on http://localhost:" + cfg.Port)

	// Create a channel to receive OS signals
	// Notify this channel when Ctrl+C or SIGTERM is received
	// Block until a shutdown signal arrives

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	slog.Info("Shutting down server...")

	// Give existing requests up to 5 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	// Stops accepting new requests and waits for active ones.
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Shutdown failed", "error", err)
		os.Exit(1)
	}

	slog.Info("Server gracefully shutdown")
}

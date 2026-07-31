package middleware

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler{

	return http.HandlerFunc(

		func(w http.ResponseWriter,r *http.Request){
			// logger:=slog.New(
			// 	// slog.NewTextHandler(os.Stdout,nil), // Easy for humans to read.
			// 	slog.NewJSONHandler(os.Stdout,nil), // Perfect for:Elasticsearch Loki Datadog Cloud Logging
			// )

			 // // Log to a specific file
 			// logFile, _ := os.Create("api.log")
			// logger:= slog.New(slog.NewJSONHandler(logFile,&slog.HandlerOptions{AddSource:true}))
			logger:= slog.New(slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{AddSource:true}))
			moddlewareLogger:=logger.With("GROUP","MIDDLEWARE") // to GROUP logs

			start := time.Now()
			next.ServeHTTP(w,r) // passes control to the next middleware or final handler.
			moddlewareLogger.Info(
				"User created",
				"user_id", 10,
				"email", "alice@example.com",
				"duration", time.Since(start),
			)
	})
}
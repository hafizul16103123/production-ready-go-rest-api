package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler{

	return http.HandlerFunc(

		func(w http.ResponseWriter,r *http.Request){
			moddlewareLogger:=slog.With("GROUP","MIDDLEWARE") // to GROUP logs
			
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
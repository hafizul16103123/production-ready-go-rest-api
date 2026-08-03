package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/auth"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/config"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/response"
)


type ContextKey string
const UserContextKey ContextKey = "user"

func Protect(fn http.HandlerFunc, roles ...string) http.Handler {

	var handler http.Handler = fn

	for _, role := range roles {
		handler = RequireRole(role)(handler)
	}

	return JWT(config.Get().JWTSecret)(handler)
}

func JWT(secret string) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			header := r.Header.Get("Authorization")

			if header == "" {

				response.Error(
					w,
					http.StatusUnauthorized,
					"missing authorization header",
				)

				return
			}

			parts := strings.Split(header, " ")

			if len(parts) != 2 || parts[0] != "Bearer" {

				response.Error(
					w,
					http.StatusUnauthorized,
					"invalid authorization header",
				)

				return
			}

			claims, err := auth.ParseToken(
				parts[1],
				secret,
			)

			if err != nil {

				response.Error(
					w,
					http.StatusUnauthorized,
					"invalid token",
				)

				return
			}

			ctx := context.WithValue(

				r.Context(),

				UserContextKey,

				claims,

			)

			next.ServeHTTP(

				w,

				r.WithContext(ctx),

			)

		})

	}

}
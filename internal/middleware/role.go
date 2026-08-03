package middleware

import (
	"net/http"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/auth"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/response"
)
func RequireRole(

	role string,

) func(http.Handler)http.Handler{

	return func(next http.Handler)http.Handler{

		return http.HandlerFunc(func(

			w http.ResponseWriter,

			r *http.Request,

		){

			claims:=r.Context().
				Value(UserContextKey).
				(*auth.Claims)

			if claims.Role!=role{

				response.Error(

					w,

					http.StatusForbidden,

					"permission denied",

				)

				return
			}

			next.ServeHTTP(w,r)

		})

	}
}
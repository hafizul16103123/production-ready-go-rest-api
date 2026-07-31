package middleware

import (
	"net/http"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/response"
)

func AuthMiddleware(next http.Handler) http.Handler{

	return http.HandlerFunc(
		func(w http.ResponseWriter,r *http.Request){

			token:= r.Header.Get("Authorization")
			if token==""{
				response.Error(w,http.StatusUnauthorized,"User is Unauthenticate")
				return
			}

			next.ServeHTTP(w,r)
		})
}
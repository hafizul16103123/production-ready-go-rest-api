package middleware

import (
	"fmt"
	"net/http"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/response"
)

func RecoveryMiddleware(next http.Handler) http.Handler{

	return http.HandlerFunc(

		func(w http.ResponseWriter,r *http.Request){
			
			defer func(){ // recover() it only works inside a deferred function.
				if err:=recover();err!=nil{
					fmt.Println(err)
					response.Error(w,http.StatusInternalServerError,"Internal Server Error")
				}
			}()

			next.ServeHTTP(w,r)
		})
}
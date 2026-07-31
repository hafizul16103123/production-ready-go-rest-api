package handler

import (
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/response"
	"log"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Method:", r.Method)
	log.Println("Remote:", r.RemoteAddr)
	log.Println("Host:", r.Host)
	log.Println("HeaderValue:", r.Header.Get("Authorization"))
	w.Header().Set("Content-Type", "application/json")

	// GET /users/{id} used as /users/25
	log.Println("Path:", r.URL.Path)              // /users
	log.Println("Path value:", r.PathValue("id")) // 25

	// GET /users?page=2&limit=105
	log.Println("Page:", r.URL.Query().Get("page"))
	log.Println("limit:", r.URL.Query().Get("limit"))

	// --- remaining *http.Request methods ---
	log.Println("Context:", r.Context())
	log.Println("ProtoAtLeast(1,0):", r.ProtoAtLeast(1, 0))
	log.Println("UserAgent:", r.UserAgent())
	log.Println("Referer:", r.Referer())
	log.Println("Cookies:", r.Cookies())

	if cookie, err := r.Cookie("session_id"); err == nil {
		log.Println("Cookie session_id:", cookie.Value)
	} else {
		log.Println("Cookie session_id:", err)
	}

	if username, password, ok := r.BasicAuth(); ok {
		log.Println("BasicAuth:", username, password)
	} else {
		log.Println("BasicAuth: not provided")
	}

	user := User{
		Id:   10,
		Name: "Hafiz",
	}

	response.WriteJSON(w, http.StatusOK, user)
}

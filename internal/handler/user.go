package handler

/*

Struct সাধারণত Pointer হিসেবে Inject করা হয়।

UserRepository একটি Interface। Interface নিজেই Reference-এর মতো আচরণ করে।
Interface নিজেই = Type + Pointer to Data

UserService is a concrete struct, so we inject it as a pointer to avoid copying and to share the same instance.
UserRepository is an interface, and interfaces already hold a reference to the concrete implementation,
so using a pointer to an interface is unnecessary and considered bad practice.

১. UserHandler → *service.UserService (pointer)
UserService একটা concrete struct। Struct কে ভ্যালু হিসেবে রাখলে প্রতিবার কপি হয়ে যায় (নতুন মেমরি অ্যালোকেশন)। Pointer রাখলে:

কপি এড়ানো যায় (efficient)
সব জায়গায় একই instance শেয়ার হয় (state share করতে পারবেন ভবিষ্যতে দরকার হলে)
২. UserService → repository.UserRepository (pointer না, সরাসরি interface)
Interface নিজেই ভেতরে ভেতরে দুইটা জিনিস বহন করে — (type, pointer-to-data)। মানে interface variable টা ইতিমধ্যেই একটা reference-এর মতো কাজ করে। তাই *repository.UserRepository লেখাটা:

অপ্রয়োজনীয় (redundant indirection)
Go community-তে এটাকে bad practice / code smell ধরা হয়
৩. MemoryRepository → users []model.User (pointer না, slice)
Slice-ও Go তে একটা reference type (ভেতরে pointer to underlying array + len + cap থাকে)। তাই এখানেও আলাদা pointer দরকার নেই।

সারমর্ম: Go তে pointer শুধু তখনই দরকার যখন underlying type হলো একটা concrete value (struct, array, int ইত্যাদি)
যেটা কপি হলে সমস্যা হবে বা shared state দরকার। কিন্তু interface, slice, map, channel — এগুলো নিজেরাই reference semantics নিয়ে আসে,
তাই এদের সামনে * বসানো ভুল/অপ্রয়োজনীয়।

*/

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/model"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/service"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/response"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(
	service *service.UserService,
) *UserHandler {

	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	time.Sleep(2 * time.Millisecond) // Simulate a delay for demonstration purposes

	users, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	response.WriteJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetUser(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, _ := strconv.Atoi(r.PathValue("id"))

	user, found := h.service.GetByID(id)

	if !found {

		http.Error(w, "User not found", http.StatusNotFound)

		return
	}

	response.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) CreateUser(
	w http.ResponseWriter,
	r *http.Request,
) {

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)

		return
	}

	user = h.service.Create(user)

	w.WriteHeader(http.StatusCreated)

	response.WriteJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, _ := strconv.Atoi(r.PathValue("id"))

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)

		return
	}

	updated, ok := h.service.Update(id, user)

	if !ok {

		http.Error(w, "User not found", http.StatusNotFound)

		return
	}

	response.WriteJSON(w, http.StatusOK, updated)
}

func (h *UserHandler) DeleteUser(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, _ := strconv.Atoi(r.PathValue("id"))

	ok := h.service.Delete(id)

	if !ok {

		http.Error(w, "User not found", http.StatusNotFound)

		return
	}

	response.WriteJSON(w, http.StatusNoContent, nil)
}

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
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/auth"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/dto"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/middleware"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/model"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/repository"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/response"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/service"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/validator"
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

// Register routes for user-related endpoints
func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /users", middleware.Protect(h.GetUsers))
	mux.Handle("GET /users/{id}", middleware.Protect(h.GetUser))
	mux.Handle("POST /users", middleware.Protect(h.CreateUser))
	mux.Handle("PUT /users/{id}", middleware.Protect(h.UpdateUser))
	mux.Handle("DELETE /users/{id}", middleware.Protect(h.DeleteUser))
}

func (h *UserHandler) CreateUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()
	var input dto.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {

		response.Error(w, http.StatusBadRequest, "Invalid JSON")

		return
	}
	fmt.Println(1)
	fmt.Println("input:", input)

	if err := validator.Validate.Struct(input); err != nil {

		response.ValidationError(
			w,
			err,
		)

		return
	}
	fmt.Println(2)
	user, err := h.service.Create(ctx,input.ToModel())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, _ := strconv.Atoi(r.PathValue("id"))

	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {

		response.Error(w, http.StatusBadRequest, "Invalid JSON")

		return
	}

	updated, err := h.service.Update(r.Context(), id, user)

	if errors.Is(err, repository.ErrNotFound) {

		response.Error(w, http.StatusNotFound, "User not found")

		return
	}

	if err != nil {

		response.Error(w, http.StatusInternalServerError, "Failed to update user")

		return
	}

	response.Success(w, http.StatusOK, updated)
}
func (h *UserHandler) GetUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	claims := r.Context().
		Value(middleware.UserContextKey).
		(*auth.Claims)
	fmt.Println(claims.Email)

	users, err := h.service.GetAll(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}
	response.Success(w, http.StatusOK, users)
}

func (h *UserHandler) GetUser(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, _ := strconv.Atoi(r.PathValue("id"))

	user, err := h.service.GetByID(r.Context(), id)

	if errors.Is(err, repository.ErrNotFound) {

		response.Error(w, http.StatusNotFound, "User not found")

		return
	}

	if err != nil {

		response.Error(w, http.StatusInternalServerError, "Failed to retrieve user")

		return
	}

	response.Success(w, http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, _ := strconv.Atoi(r.PathValue("id"))

	err := h.service.Delete(r.Context(), id)

	if errors.Is(err, repository.ErrNotFound) {

		response.Error(w, http.StatusNotFound, "User not found")

		return
	}

	if err != nil {

		response.Error(w, http.StatusInternalServerError, "Failed to delete user")

		return
	}

	response.Success(w, http.StatusNoContent, nil)
}

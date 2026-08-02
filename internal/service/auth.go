package service

import (
	"context"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/auth"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/model"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/repository"
)

type AuthService struct{
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (authService *AuthService) Register(ctx context.Context, user model.User) (model.User, error) {
	hashedPassword, err := auth.HashPassword(user.PasswordHash)
	if err != nil {
		return model.User{}, err
	}
	user.PasswordHash = string(hashedPassword)
	return authService.repo.Create(ctx,user)
}

func (authService *AuthService) Login(user model.User) (model.User, error) {
	err := auth.CheckPasswordHash(user.PasswordHash, user.PasswordHash)
	if err != nil {
		return model.User{}, err
	}
	return authService.repo.GetByEmail(user.Email)
}
package service

import (
	"context"
	"fmt"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/auth"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/config"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/dto"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/model"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/repository"
)

type AuthService struct {
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
	return authService.repo.Create(ctx, user)
}

func (authService *AuthService) Login(ctx context.Context, authDTO dto.LoginDTO) (model.User, string, error) {
	user, err := authService.repo.GetByEmail(ctx, authDTO.Email)
	fmt.Println("user:", user)
	if err != nil {
		return model.User{}, "", err
	}
	if err := auth.CheckPassword(authDTO.Password, user.PasswordHash); err != nil {
		return model.User{}, "", err
	}
	token, err := auth.GenerateToken(
		user.Id,
		user.Email,
		user.Role,
		config.Get().JWTSecret,
	)
	if err != nil {
		return model.User{}, "", err
	}
	return user, token, nil
}

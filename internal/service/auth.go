package service

import (
	"context"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/model"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{
	repo repository.UserRepository
}

func (authService *AuthService) Register(ctx context.Context, user model.User) (model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
	[]byte(user.Password),
	bcrypt.DefaultCost,
)
	if err != nil {
		return model.User{}, err
	}
	user.Password = string(hashedPassword)
	return authService.repo.Create(ctx,user)
}

// func (authService *AuthService) Login(user model.User) (model.User, error) {
// 	err := bcrypt.CompareHashAndPassword(
// 		[]byte(user.PasswordHash),
// 		[]byte(password),
// 	)
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	return authService.repo.GetByEmail(user.Email)
// }
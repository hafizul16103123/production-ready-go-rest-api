package service

import (
	"context"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/model"
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}

}

func (s *UserService) GetAll(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id int) (model.User, error) {

	return s.repo.GetByID(ctx, id)
}

func (s *UserService) Create(ctx context.Context, user model.User) (model.User, error) {

	return s.repo.Create(ctx, user)
}

func (s *UserService) Update(ctx context.Context, id int, user model.User) (model.User, error) {

	return s.repo.Update(ctx, id, user)
}

func (s *UserService) Delete(ctx context.Context, id int) error {

	return s.repo.Delete(ctx, id)
}

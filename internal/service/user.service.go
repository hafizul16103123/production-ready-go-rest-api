package service

import (
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

func (s *UserService) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetByID(id int) (model.User, bool) {

	return s.repo.GetUserById(id)
}

func (s *UserService) Create(user model.User) model.User {

	return s.repo.Create(user)
}

func (s *UserService) Update(id int, user model.User) (model.User, bool) {

	return s.repo.Update(id, user)
}

func (s *UserService) Delete(id int) bool {

	return s.repo.Delete(id)
}

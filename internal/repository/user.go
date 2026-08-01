package repository

import (
	"github.com/hafizul16103123/production-ready-go-rest-api/internal/model"
)

type UserRepository interface {
	Create(user model.User) (model.User, error)
	Update(id int, user model.User) (model.User, bool)
	Delete(id int) bool
	GetAll() ([]model.User, error)
	GetUserById(id int) (model.User, bool)
}

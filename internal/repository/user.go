package repository

import (
	"context"
	"errors"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/model"
)

var ErrNotFound = errors.New("user not found")

type UserRepository interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	Update(ctx context.Context, id int, user model.User) (model.User, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id int) (model.User, error)
}

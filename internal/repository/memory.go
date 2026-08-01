package repository

import (
	"context"
	"fmt"

	"github.com/hafizul16103123/production-ready-go-rest-api/internal/model"
)

type MemoryRepository struct {
	users []model.User
}

func NewMemoryRepository() *MemoryRepository {

	return &MemoryRepository{
		users: []model.User{
			{
				Id:    1,
				Name:  "Alice",
				Email: "alice@gmail.com",
			},
			{
				Id:    2,
				Name:  "Bob",
				Email: "bob@gmail.com",
			},
		},
	}
}

func (m *MemoryRepository) GetAll(ctx context.Context) ([]model.User, error) {
	fmt.Println("GetAll called from MemoryRepository")
	return m.users, nil
}

func (m *MemoryRepository) GetByID(ctx context.Context, Id int) (model.User, error) {

	for _, user := range m.users {

		if user.Id == Id {

			return user, nil

		}
	}

	return model.User{}, ErrNotFound
}

func (m *MemoryRepository) Create(ctx context.Context,user model.User) (model.User, error) {
	user.Id = len(m.users) + 1

	m.users = append(m.users, user)

	return user, nil
}

func (m *MemoryRepository) Update(ctx context.Context, Id int, user model.User) (model.User, error) {

	for i := range m.users {

		if m.users[i].Id == Id {

			user.Id = Id

			m.users[i] = user

			return user, nil
		}
	}

	return model.User{}, ErrNotFound
}

func (m *MemoryRepository) Delete(ctx context.Context, Id int) error {

	for i := range m.users {

		if m.users[i].Id == Id {

			m.users = append(
				m.users[:i],
				m.users[i+1:]...,
			)

			return nil
		}
	}

	return ErrNotFound
}

package repository

import (
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

func (m *MemoryRepository) GetAll() ([]model.User, error) {
	fmt.Println("GetAll called from MemoryRepository")
	return m.users, nil
}

func (m *MemoryRepository) GetUserById(Id int) (model.User, bool) {

	for _, user := range m.users {

		if user.Id == Id {

			return user, true

		}
	}

	return model.User{}, false
}

func (m *MemoryRepository) Create(user model.User) model.User {

	user.Id = len(m.users) + 1

	m.users = append(m.users, user)

	return user
}

func (m *MemoryRepository) Update(Id int, user model.User) (model.User, bool) {

	for i := range m.users {

		if m.users[i].Id == Id {

			user.Id = Id

			m.users[i] = user

			return user, true
		}
	}

	return model.User{}, false
}

func (m *MemoryRepository) Delete(Id int) bool {

	for i := range m.users {

		if m.users[i].Id == Id {

			m.users = append(
				m.users[:i],
				m.users[i+1:]...,
			)

			return true
		}
	}

	return false
}

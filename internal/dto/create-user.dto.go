package dto

import "github.com/hafizul16103123/production-ready-go-rest-api/internal/model"

type CreateUserDTO struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,gte=0,lte=120"`
	Password string `json:"password" validate:"required,min=8"`
}

func (d CreateUserDTO) ToModel() model.User {
	return model.User{
		Name:  d.Name,
		Email: d.Email,
		Age:   d.Age,
		Password: d.Password,
	}
}
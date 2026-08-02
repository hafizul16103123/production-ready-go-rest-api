package dto

import "github.com/hafizul16103123/production-ready-go-rest-api/internal/model"

type UserResponseDTO struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func UserResponseFromModel(user model.User) UserResponseDTO {
	return UserResponseDTO{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
	}
}

type LoginResponseDTO struct {
	User  UserResponseDTO `json:"user"`
	Token string          `json:"token"`
}

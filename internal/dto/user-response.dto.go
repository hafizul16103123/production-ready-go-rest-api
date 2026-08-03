package dto

import "github.com/hafizul16103123/production-ready-go-rest-api/internal/model"

type UserResponseDTO struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	Role  string `json:"role"`
}

func UserResponseFromModel(user model.User) UserResponseDTO {
	return UserResponseDTO{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Age:   user.Age,
		Role:  user.Role,
	}
}

func UserResponsesFromModel(users []model.User) []UserResponseDTO {
	responses := make([]UserResponseDTO, len(users))

	for i, user := range users {
		responses[i] = UserResponseFromModel(user)
	}

	return responses
}

type LoginResponseDTO struct {
	User  UserResponseDTO `json:"user"`
	Token string          `json:"token"`
}

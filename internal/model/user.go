package model

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Age          int    `json:"age"`
	PasswordHash string `json:"passwordHash"`
	Role		 string `json:"role" validate:" required oneof=admin user"`
}

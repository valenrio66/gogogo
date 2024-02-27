package user

import (
	"gogogo/internal/repository"
	v1 "gogogo/internal/service/v1"
)

type userHandler struct {
	User repository.User
}

func NewUserHandler(user repository.User) v1.User {
	return &userHandler{
		User: user,
	}
}

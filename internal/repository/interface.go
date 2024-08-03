package repository

import (
	"context"
	"gogogo/internal/model"
)

type User interface {
	GetDataUser(ctx context.Context) (user []model.User, err error)
	CreateUser(ctx context.Context, user model.User) (createdUser model.User, err error)
}

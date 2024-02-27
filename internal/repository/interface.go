package repository

import (
	"context"
	"gogogo/internal/model"
)

type User interface {
	GetDataUser(ctx context.Context) (user model.User, err error)
}

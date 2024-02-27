package user

import (
	"context"
	"gogogo/internal/repository"
	"gorm.io/gorm"
)

type UserItab struct {
	Gorm *gorm.DB
}

func UserTable(db *gorm.DB) repository.User {
	return &UserItab{Gorm: db.Table("user").
		Session(&gorm.Session{}).
		WithContext(context.Background())}
}

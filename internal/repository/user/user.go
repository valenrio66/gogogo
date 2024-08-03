package user

import (
	"context"
	"gogogo/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (db *UserItab) GetDataUser(ctx context.Context) (user []model.User, err error) {
	grm := db.Gorm
	err = grm.WithContext(ctx).Model(&model.User{}).
		Find(&user).
		Error
	return
}

func (db *UserItab) CreateUser(ctx context.Context, user model.User) (createdUser model.User, err error) {
	grm := db.Gorm

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return createdUser, err
	}

	user.Password = string(hashedPassword)

	// Save user
	err = grm.WithContext(ctx).Create(&user).Error
	if err != nil {
		return createdUser, err
	}

	createdUser = user
	return createdUser, nil
}

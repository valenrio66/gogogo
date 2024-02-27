package user

import (
	"context"
	"gogogo/internal/model"
)

func (db *UserItab) GetDataUser(ctx context.Context) (user model.User, err error) {
	grm := db.Gorm
	err = grm.WithContext(ctx).Model(&model.User{}).
		Find(&user).
		Error
	return
}

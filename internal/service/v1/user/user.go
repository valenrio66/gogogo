package user

import (
	"fmt"
	fibhelp "github.com/aiteung/athelper/fiber"
	"github.com/gofiber/fiber/v2"
	"gogogo/internal/model"
)

func (u *userHandler) GetAllUser(ctx *fiber.Ctx) (err error) {
	dataUser, err := u.User.GetDataUser(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data User tidak Ditemukan")
	}

	err = fibhelp.ReturnData[model.User]{
		Code:    fiber.StatusOK,
		Success: true,
		Status:  fmt.Sprintf("Berhasil Get Data"),
		Data:    dataUser,
	}.WriteResponseBody(ctx)

	return
}

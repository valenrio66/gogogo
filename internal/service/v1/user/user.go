package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valenrio66/gofuck"
)

func (u *userHandler) GetAllUser(ctx *fiber.Ctx) (err error) {
	dataUser, err := u.User.GetDataUser(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data User tidak Ditemukan")
	}

	return gofuck.GoddamnResponse(ctx, fiber.StatusOK, true, "Berhasil Get Data", dataUser)
}

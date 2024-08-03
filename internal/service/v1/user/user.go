package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valenrio66/gofuck"
	"gogogo/internal/model"
)

func (u *userHandler) GetAllUser(ctx *fiber.Ctx) (err error) {
	dataUser, err := u.User.GetDataUser(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data User tidak Ditemukan")
	}

	return gofuck.GoddamnResponse(ctx, fiber.StatusOK, true, "Berhasil Get Data", dataUser)
}

func (u *userHandler) CreateUser(ctx *fiber.Ctx) (err error) {
	var user model.User
	if err := gofuck.ParseFuckingBody(ctx, &user); err != nil {
		return err
	}

	createdUser, err := u.User.CreateUser(ctx.Context(), user)
	if err != nil {
		return gofuck.FuckTheErrorResponse(ctx, fiber.StatusInternalServerError, "Gagal Membuat User")
	}

	return gofuck.GoddamnResponse(ctx, fiber.StatusOK, true, "Berhasil Create Data", createdUser)
}

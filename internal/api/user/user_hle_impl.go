package user

import (
	"chat-app-fiber/internal/service/userServ"
	"github.com/gofiber/fiber/v2"
)

type userApiHandler struct {
	userServ userServ.UserService
}

func NewUserApiHandler(userServ userServ.UserService) UserApiHandler {
	return &userApiHandler{
		userServ: userServ,
	}
}

func (u userApiHandler) AuthenticationUser(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u userApiHandler) RegisterUser(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

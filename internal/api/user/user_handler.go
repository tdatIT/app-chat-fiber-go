package user

import "github.com/gofiber/fiber/v2"

type UserApiHandler interface {
	AuthenticationUser(ctx *fiber.Ctx) error
	RegisterUser(ctx *fiber.Ctx) error
}

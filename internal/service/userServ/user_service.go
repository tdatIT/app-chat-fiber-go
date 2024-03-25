package userServ

import "chat-app-fiber/internal/domain/dto/user"

type UserService interface {
	Authentication(req *user.LoginRequest) (*user.LoginResponse, error)
	RegisterUser(req *user.CreateUserRequest) (*user.CreateUserResponse, error)
}

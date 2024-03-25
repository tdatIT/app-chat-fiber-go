package userRepos

import "chat-app-fiber/internal/domain/entities"

type UserRepository interface {
	// queries
	FindAll(size int, page int) ([]*entities.User, int, error)
	FindByID(id string) (*entities.User, error)
	FindByUsername(username string) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)

	// commands
	CreateUser(user *entities.User) (entities.User, error)
	UpdateUserInformation(user *entities.User) error
	InActiveUser(id string) error
}

package userRepos

import (
	"chat-app-fiber/internal/domain/entities"
	"chat-app-fiber/internal/domain/repositories"
	"chat-app-fiber/internal/utils/db/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(dbClient *mongodb.MongoClient) repositories.UserRepository {
	col := dbClient.GetDB().Collection("users")
	return &userRepository{userCollection: col}
}

func (u userRepository) FindAll(size int, page int) ([]*entities.User, int, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindByID(id string) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindByUsername(username string) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindByEmail(email string) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) CreateUser(user *entities.User) (entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) UpdateUserInformation(user *entities.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) InActiveUser(id string) error {
	//TODO implement me
	panic("implement me")
}

package userRepos

import (
	"chat-app-fiber/internal/domain/entities"
	"chat-app-fiber/internal/domain/repositories"
	"github.com/stretchr/testify/mock"
)

type userReposMock struct {
	mock *mock.Mock
}

func NewUserReposMock() repositories.UserRepository {
	return &userReposMock{}
}

func (u userReposMock) FindAll(size int, page int) ([]*entities.User, int, error) {
	return u.mock.Called(size, page).Get(0).([]*entities.User), u.mock.Called(size, page).Get(1).(int), u.mock.Called(size, page).Error(2)
}

func (u userReposMock) FindByID(id string) (*entities.User, error) {
	return u.mock.Called(id).Get(0).(*entities.User), u.mock.Called(id).Error(1)
}

func (u userReposMock) FindByUsername(username string) (*entities.User, error) {
	return u.mock.Called(username).Get(0).(*entities.User), u.mock.Called(username).Error(1)
}

func (u userReposMock) FindByEmail(email string) (*entities.User, error) {
	return u.mock.Called(email).Get(0).(*entities.User), u.mock.Called(email).Error(1)
}

func (u userReposMock) CreateUser(user *entities.User) (entities.User, error) {
	return u.mock.Called(user).Get(0).(entities.User), u.mock.Called(user).Error(1)
}

func (u userReposMock) UpdateUserInformation(user *entities.User) error {
	return u.mock.Called(user).Error(0)
}

func (u userReposMock) InActiveUser(id string) error {
	return u.mock.Called(id).Error(0)
}

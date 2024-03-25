package repositories

import (
	"chat-app-fiber/internal/domain/repositories/userRepos"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	userRepos.NewUserRepository,
	userRepos.NewUserReposMock,
)

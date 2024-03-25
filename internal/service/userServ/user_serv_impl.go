package userServ

import (
	"chat-app-fiber/config"
	"chat-app-fiber/internal/domain/dto/user"
	"chat-app-fiber/internal/domain/entities"
	"chat-app-fiber/internal/domain/repositories/userRepos"
	"chat-app-fiber/internal/utils/errorUtils"
	"chat-app-fiber/internal/utils/mapper"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type userServiceImpl struct {
	cfg      *config.AppConfig
	userRepo userRepos.UserRepository
}

func NewUserService(config *config.AppConfig, userRepo userRepos.UserRepository) UserService {
	return &userServiceImpl{config, userRepo}
}

func (u userServiceImpl) Authentication(req *user.LoginRequest) (*user.LoginResponse, error) {
	var dao *entities.User
	var err error

	if IsEmail(req.Username) {
		dao, err = u.userRepo.FindByEmail(req.Username)
	} else {
		dao, err = u.userRepo.FindByUsername(req.Username)
	}

	if err != nil {
		log.Error(err)
		return nil, err
	}

	if !CheckPasswordHash(req.Password, dao.Password) {
		return nil, errorUtils.ErrInvalidCredentials
	}

	exprTime := time.Now().Add(time.Hour * 24 * 7)

	claims := jwt.MapClaims{
		"id":       dao.ID,
		"email":    dao.Email,
		"username": dao.Username,
		"avt":      dao.Avatar,
		"roles":    dao.Roles,
		"exp":      exprTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(u.cfg.SecretKey))
	if err != nil {
		return nil, errorUtils.ErrInternalServer.WithInternalError(err)
	}

	resp := user.LoginResponse{
		Token:   tokenStr,
		Expired: int(exprTime.Unix()),
	}

	return &resp, nil
}

func (u userServiceImpl) RegisterUser(req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	hashPass, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	var dao entities.User
	if err := mapper.BindingStruct(req, dao); err != nil {
		return nil, err
	}

	dao.Password = hashPass
	dao.Roles = []string{entities.ROLE_USER}
	dao.CreatedAt = time.Now()
	dao.UpdatedAt = time.Now()

	_, err = u.userRepo.CreateUser(&dao)
	if err != nil {
		return nil, err
	}

	var resp user.CreateUserResponse
	if err := mapper.BindingStruct(dao, resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

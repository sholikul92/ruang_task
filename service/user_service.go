package service

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
	"a21hc3NpZ25tZW50/utils"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserServiceInterface interface {
	Register(payload *model.Register) error
	Login(payload *model.Login) (token *string, userId int, err error)
}

type userService struct {
	userRepo repository.UserRepoInterface
}

func NewUserService(userRepo repository.UserRepoInterface) *userService {
	return &userService{userRepo: userRepo}
}

func (us *userService) Register(payload *model.Register) error {
	userExists, _ := us.userRepo.FindUser(payload.Username)

	if userExists != nil && userExists.Username == payload.Username {
		return errors.New("username already exists")
	}

	hashPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	var user = model.User{
		Username: payload.Username,
		Password: hashPassword,
	}

	if err := us.userRepo.CreateUser(&user); err != nil {
		return err
	}

	return nil
}

func (us *userService) Login(payload *model.Login) (token *string, userID int, err error) {
	user, err := us.userRepo.FindUser(payload.Username)
	if err != nil {
		return nil, 0, errors.New("user not registered")
	}

	if err := utils.ComparePassword(user.Password, payload.Password); err != nil {
		return nil, 0, errors.New("wrong password!")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &model.Claims{
		UserId:   int(user.ID),
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(model.JwtKey)
	if err != nil {
		return nil, 0, err
	}

	return &tokenString, claims.UserId, nil
}

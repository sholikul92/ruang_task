package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepoInterface interface {
	CreateUser(user *model.User) error
	FindUser(username string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (ur *userRepo) CreateUser(user *model.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) FindUser(username string) (*model.User, error) {
	var user model.User
	if err := ur.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

package services

import (
	"BukaLelang/app/features/users"
	"BukaLelang/helper"
	"errors"
	"fmt"

	"github.com/labstack/gommon/email"
)

type UserService struct {
	m users.Repository
}

func New(r users.Repository) users.Repository {
	return &UserService{m:r}
}

func (us *UserService) Register(newUser users.Core) error {
	_, err := us.m.Register(newUser)
    if err != nil {
		return errors.New("Failed to register user")
	}
	return nil
} 

func (us *UserService) Login(email string, password string) (users.Core, error) {
	user, err := us.m.Login(email, password)
	if err != nil {
		return users.Core{},err
	}
	return user, nil
}

func (us *UserService) GetProfile(id uint) (users.Core, error) {
	user, err := us.m.GetProfile(id)
	if err != nil {
		return users.Core{},err
	}
	return user, nil
} 

func (us *UserService) UpdateProfile(id uint, updateUser users.Core) error {
		hashedPassword, err := helper.HashedPassword(updateUser.Password)
		if err != nil {
				return fmt.Errorf("failed to hash password: %v", err)
		}
		updateUser := users.Core{
			Username:  updateUser.Username,
			Email:     updateUser.Email,
			Phone:     updateUser.Phone,
			Password:  string(hashedPassword),
			Image:     updateUser.Image,
		}
		if err := us.m.UpdateProfile(id, updateUser); err != nil {
				return fmt.Errorf("Error while updating %d: %v", id, err)
		}
		return nil 
}
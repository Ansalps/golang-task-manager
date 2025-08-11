package service

import (
	"errors"

	"github.com/Ansalps/golang-task-manager/internal/model"
	repointerface "github.com/Ansalps/golang-task-manager/internal/repository/interfaces"
	serviceinterface "github.com/Ansalps/golang-task-manager/internal/service/interfaces"
	"gorm.io/gorm"
)

type userService struct {
	userRepo repointerface.UserRepository // interface
}

func NewUserService(userRepo repointerface.UserRepository) serviceinterface.UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user *model.User) error {
	_, err := s.userRepo.FindByEmail(user.Email)
	if err != gorm.ErrRecordNotFound {
		return errors.New("user already Exists")
	}
	return s.userRepo.Create(user)
}

func (s *userService) FindByID(id uint) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (c *userService) UserLogin(user *model.UserLogin) (*model.User, error) {
	User, err := c.userRepo.FindByEmail(user.Email)
	if err != nil {
		return nil, errors.New("email or password is incorrect")
	}
	if user.Password != User.Password {
		return nil, errors.New("email or password is incorrect")
	}
	return User, nil
}

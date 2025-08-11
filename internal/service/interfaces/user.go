package interfaces

import (
	"github.com/Ansalps/golang-task-manager/internal/model"
)

type UserService interface {
	FindByID(id uint) (*model.User, error)
	CreateUser(user *model.User) error
	UserLogin(user *model.UserLogin) (*model.User, error)
}

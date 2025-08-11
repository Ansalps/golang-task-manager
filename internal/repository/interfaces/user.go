package interfaces

import (
	"github.com/Ansalps/golang-task-manager/internal/model"
	
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	
}

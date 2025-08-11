package repository

import (
	"errors"

	"github.com/Ansalps/golang-task-manager/internal/model"
	"github.com/Ansalps/golang-task-manager/internal/repository/interfaces"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{DB: db}
}

// Create a new user
func (r *userRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

// Get user by email
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

var ErrUserNotFound = errors.New("user not found")

func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.DB.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

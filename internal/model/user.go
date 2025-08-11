package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primarykey" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	FirstName string    `json:"first_name" validate:"required,excludesall= "`
	LastName  string    `json:"last_name" validate:"required,nameOrInitials"`
	Email     string    `json:"email"  validate:"required,email" gorm:"unique"`
	Password  string    `json:"password" validate:"required,password,min=8"`
	ConfirmPassword string `gorm:"-" json:"confirm_password"`
	Phone     string    `json:"phone" validate:"required,numeric,len=10"`
}
type UserLogin struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}


package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint           `gorm:"primarykey" json:"-"`
	UserID      uint           `gorm:"not null;index" json:"-"`
	User        User           `gorm:"foreignKey:UserID;references:ID" json:"-" validate:"-"`
	Title       string         `gorm:"not null" json:"title" validate:"required,no_leading_trailing_spaces"`
	Description string         `json:"description" validate:"required,no_leading_trailing_spaces"`
	Status      string         `gorm:"default:'todo';check:status IN ('todo', 'inprogress', 'done')" validate:"required,oneof=todo inprogress done" json:"status"` // todo, inprogress, done
	DueAt       string     `gorm:"type:date" json:"due_at" validate:"required"`
	IsPublic    bool           `gorm:"default:false" json:"is_public"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

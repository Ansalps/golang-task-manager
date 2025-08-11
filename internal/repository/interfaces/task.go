package interfaces

import (
	"time"

	"github.com/Ansalps/golang-task-manager/internal/model"
)

type TaskRepository interface {
	CreateTask(task *model.Task) error
	GetTaskByID(id, userid uint) (*model.Task, error)
	GetUserTasks(userID uint, page, limit int, status string, dueAt *time.Time) ([]model.Task, error)
	UpdateTask(task *model.Task, taskid, userid uint) error
	DeleteTask(id uint, userID uint) error
	GetPublicTasks() ([]model.Task, error)
}

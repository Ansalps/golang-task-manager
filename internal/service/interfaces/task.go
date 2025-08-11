package interfaces

import "github.com/Ansalps/golang-task-manager/internal/model"

type TaskService interface {
	CreateTask(task *model.Task) error
	GetTaskByID(id,userid uint) (*model.Task, error)
	GetAllTasks(userID uint,pageParam,limitParam string ,status,dueAtStr string) ([]model.Task, error)
	UpdateTask(task *model.Task,taskid,userid uint) error
	DeleteTask(taskid,userid uint) error
	GetPublicTasks() ([]model.Task,error)
}

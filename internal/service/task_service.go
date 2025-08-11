package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/Ansalps/golang-task-manager/internal/model"
	repointerface "github.com/Ansalps/golang-task-manager/internal/repository/interfaces"
	serviceinterface "github.com/Ansalps/golang-task-manager/internal/service/interfaces"
)

type taskService struct {
	taskRepo repointerface.TaskRepository //interface
}

func NewTaskService(taskRepo repointerface.TaskRepository) serviceinterface.TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}
func (s *taskService) CreateTask(task *model.Task) error {
	return s.taskRepo.CreateTask(task)
}

func (s *taskService) GetTaskByID(id, userid uint) (*model.Task, error) {
	return s.taskRepo.GetTaskByID(id, userid)
}

func (s *taskService) GetAllTasks(userID uint, pageParam, limitParam string, status, dueAtStr string) ([]model.Task, error) {
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}
	var dueAt *time.Time
	if dueAtStr != "" {
		parsedDate, err := time.Parse("2006-01-02", dueAtStr)
		if err != nil {
			return nil, errors.New("Invalid due_at format, expected YYYY-MM-DD")
		}
		dueAt = &parsedDate
	}
	// Limit sanity check
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}

	// Page sanity check
	if page < 1 {
		page = 1
	}

	if status != "" && status != "todo" && status != "inprogress" && status != "done" {
		return nil, errors.New("invalid status value")
	}
	return s.taskRepo.GetUserTasks(userID, page, limit, status, dueAt)
}

func (s *taskService) UpdateTask(task *model.Task, taskid, userid uint) error {
	return s.taskRepo.UpdateTask(task, taskid, userid)
}

func (s *taskService) DeleteTask(taskid, userid uint) error {
	return s.taskRepo.DeleteTask(taskid, userid)
}
func (s *taskService) GetPublicTasks() ([]model.Task, error) {
	return s.taskRepo.GetPublicTasks()
}

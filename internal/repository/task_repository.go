package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/Ansalps/golang-task-manager/internal/model"
	"github.com/Ansalps/golang-task-manager/internal/repository/interfaces"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) interfaces.TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task *model.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetTaskByID(id, userid uint) (*model.Task, error) {
	var task model.Task
	if err := r.db.Where("id = ? AND user_id = ?", id, userid).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) GetUserTasks(userID uint, page, limit int, status string, dueAt *time.Time) ([]model.Task, error) {

	offset := (page - 1) * limit

    query := r.db.Where("user_id = ?", userID)

    if status != "" {
        query = query.Where("status = ?", status)
    }

    if dueAt != nil {
        query = query.Where("due_at = ?", dueAt)
    }

    var tasks []model.Task
    if err := query.
        Limit(limit).
        Offset(offset).
        Order("created_at DESC").
        Find(&tasks).Error; err != nil {
        return nil, err
    }

    return tasks, nil
}

func (r *taskRepository) UpdateTask(task *model.Task, taskid, userid uint) error {
	var existing model.Task
	if err := r.db.Where("id = ? AND user_id = ?", taskid, userid).First(&existing).Error; err != nil {
		return err // returns gorm.ErrRecordNotFound if not found
	}

	if err := r.db.Model(&existing).Updates(task).Error; err != nil {
		return errors.New("failed to update specific task")
	}

	return nil
}

func (r *taskRepository) DeleteTask(id uint, userID uint) error {
	res := r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Task{})
	if res.RowsAffected == 0 {
		return errors.New("task not found or unauthorized")
	}
	return res.Error
}

func (r *taskRepository) GetPublicTasks() ([]model.Task, error) {
	var tasks []model.Task
	if err := r.db.Where("is_public = ?", true).
		Order("created_at DESC").
		Find(&tasks).Error; err != nil {
			fmt.Println("error in doing database query is",err)
		return nil, err
	}
	return tasks, nil
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/Ansalps/golang-task-manager/internal/model"
	"github.com/Ansalps/golang-task-manager/internal/service/interfaces"
	"github.com/Ansalps/golang-task-manager/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	taskService interfaces.TaskService
	userService interfaces.UserService
}

func NewTaskHandler(taskService interfaces.TaskService, userService interfaces.UserService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
		userService: userService,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	userID, err := utils.FindUserIDFromContext(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, "failed to fetch user id from context", err)
		return
	}
	// Ensure user exists
	user, err := h.userService.FindByID(userID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "User not found", err)
		return
	}

	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid task data", err)
		return
	}

	if err := utils.Validate(task); err != nil {
		utils.Error(c, http.StatusBadRequest, "validation error", err)
		return
	}

	task.UserID = user.ID
	if err := h.taskService.CreateTask(&task); err != nil {
		utils.Error(c, http.StatusInternalServerError, "Could not create task", err)
		return
	}

	utils.Created(c, "Task created successfully", task)
}

func (h *TaskHandler) GetAllTasks(c *gin.Context) {
	userID, err := utils.FindUserIDFromContext(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, "failed to fetch user id from context", err)
		return
	}
	// Ensure user exists
	user, err := h.userService.FindByID(userID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "User not found", err)
		return
	}
	// Query params
	status := c.Query("status") // empty string if not provided

	dueAtStr := c.Query("due_at") // empty string if not provided

	pageParam := c.DefaultQuery("page", "1")
	limitParam := c.DefaultQuery("limit", "10")

	// Call service/repository
	tasks, err := h.taskService.GetAllTasks(user.ID, pageParam, limitParam, status, dueAtStr)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Could not retrieve tasks", err)
		return
	}
	fmt.Println("hi.....hello.......\n", tasks)
	utils.Success(c, "Tasks retrieved successfully", tasks)
}
func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	userID, err := utils.FindUserIDFromContext(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, "failed to fetch user id from context", err)
		return
	}
	// Ensure user exists
	user, err := h.userService.FindByID(userID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "User not found", err)
		return
	}

	taskID, err := utils.StringConversion(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "invalid task id type", err)
		return
	}
	task, err := h.taskService.GetTaskByID(taskID, user.ID)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to fetch task", err)
		return
	}
	utils.Success(c, "task fetched successfully", task)
}
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	userID, err := utils.FindUserIDFromContext(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, "failed to fetch user id from context", err)
		return
	}
	// Ensure user exists
	user, err := h.userService.FindByID(userID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "User not found", err)
		return
	}

	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid task data", err)
		return
	}

	if err := utils.Validate(task); err != nil {
		utils.Error(c, http.StatusBadRequest, "validation error", err)
		return
	}

	taskID, err := utils.StringConversion(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "invalid task id type", err)
		return
	}

	err = h.taskService.UpdateTask(&task, taskID, user.ID)
	if err == gorm.ErrRecordNotFound {
		utils.Error(c, http.StatusBadRequest, "such a task do not exist for this specific user", err)
		return
	}
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Could not create task", err)
		return
	}
	utils.Success(c, "Successfully updated task", nil)
}
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	userID, err := utils.FindUserIDFromContext(c)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, "failed to fetch user id from context", err)
		return
	}
	// Ensure user exists
	user, err := h.userService.FindByID(userID)
	if err != nil {
		utils.Error(c, http.StatusNotFound, "User not found", err)
		return
	}

	taskID, err := utils.StringConversion(c.Param("id"))
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "invalid task id type", err)
		return
	}

	if err := h.taskService.DeleteTask(taskID, user.ID); err != nil {
		utils.Error(c, http.StatusInternalServerError, "Could not delete task", err)
		return
	}
	utils.Success(c, "succesfully deleted task", nil)
}

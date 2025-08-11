package handler

import (
	"fmt"
	"net/http"

	"github.com/Ansalps/golang-task-manager/internal/middleware"
	"github.com/Ansalps/golang-task-manager/internal/model"
	"github.com/Ansalps/golang-task-manager/internal/service/interfaces"
	"github.com/Ansalps/golang-task-manager/internal/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService interfaces.UserService
	taskService interfaces.TaskService
}

func NewUserHandler(userService interfaces.UserService, taskService interfaces.TaskService) *UserHandler {
	return &UserHandler{
		userService: userService,
		taskService: taskService,
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}
	if err := utils.Validate(user); err != nil {
		utils.Error(c, http.StatusBadRequest, "validation error", err)
		return
	}
	err := h.userService.CreateUser(&user)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Could not create user", err)
		return
	}

	utils.Created(c, "User created successfully", nil)
}
func (h *UserHandler) LoginUser(c *gin.Context) {
	var userlogin model.UserLogin
	if err := c.ShouldBindJSON(&userlogin); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid request data", err)
		return
	}
	if err := utils.Validate(userlogin); err != nil {
		utils.Error(c, http.StatusBadRequest, "validation error", err)
		return
	}
	User, err := h.userService.UserLogin(&userlogin)
	if err != nil {
		utils.Error(c, http.StatusUnauthorized, "enter correct information", err)
	}

	accessToken, err := middleware.GenerateToken(User.ID, User.Email, "user")
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "failed to generate jwt", err)
		return
	}
	utils.Success(c, "login successfully", accessToken)
}
func (h *UserHandler) ViewPublicTask(c *gin.Context) {
	tasks, err := h.taskService.GetPublicTasks()
	if err != nil {
		fmt.Println("tell me error in handler", err)
		utils.Error(c, http.StatusInternalServerError, "couldn't fetch public tasks", err)
		return
	}
	utils.Success(c, "fetched public tasks successfully", tasks)
}

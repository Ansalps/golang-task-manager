package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ansalps/golang-task-manager/config"
	"github.com/Ansalps/golang-task-manager/internal/handler"
	"github.com/Ansalps/golang-task-manager/internal/middleware"
	"github.com/Ansalps/golang-task-manager/internal/model"
	"github.com/Ansalps/golang-task-manager/internal/repository"
	"github.com/Ansalps/golang-task-manager/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//loading configurations
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("failed to load config file")
	}
	dbHost := viper.GetString("DB_HOST")

	// If running inside Docker, override DB_HOST
	if os.Getenv("DOCKER_ENV") == "true" {
		dbHost = "db" // service name from docker-compose
	}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		dbHost,
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)

	c.DBUrl = dbURL
	//connecting database
	db, err := gorm.Open(postgres.Open(c.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	//automigrating tables
	if err := db.AutoMigrate(&model.User{}, &model.Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	//injecting dependencies
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService, userService)

	userHandler := handler.NewUserHandler(userService, taskService)

	//creating a router
	router := gin.Default()
	//registering routes
	// auth
	router.POST("/register", userHandler.RegisterUser)
	router.POST("/login", userHandler.LoginUser)

	// public
	router.GET("/public/view-public-task", userHandler.ViewPublicTask)
	// protected
	auth := router.Group("/")
	auth.Use(middleware.JWTAuthMiddleware(c.JWTSecretKey))
	{
		auth.POST("/tasks", taskHandler.CreateTask)
		auth.GET("/tasks", taskHandler.GetAllTasks)
		auth.GET("/tasks/:id", taskHandler.GetTaskByID)
		auth.PUT("/tasks/:id", taskHandler.UpdateTask)
		auth.DELETE("/tasks/:id", taskHandler.DeleteTask)
	}
	if err := router.Run(c.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	} else {
		fmt.Println("server running successfully")
	}
}

package main

import (
	"Rest_go/internal/database"
	"Rest_go/internal/handlers"
	"Rest_go/internal/tasksService"
	"Rest_go/internal/userService"
	"Rest_go/internal/web/tasks"
	"Rest_go/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	// Проверка ошибки при автоматической миграции
	if err := database.DB.AutoMigrate(&tasksService.Task{}, &userService.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Создание репозиториев, сервисов и обработчиков
	tasksRepo := tasksService.NewTaskRepository(database.DB)
	tasksService := tasksService.NewService(tasksRepo)
	tasksHandler := handlers.NewTasksHandler(tasksService)

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Инициализация Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрация маршрутов
	tasks.RegisterHandlers(e, tasksHandler)
	users.RegisterHandlers(e, userHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

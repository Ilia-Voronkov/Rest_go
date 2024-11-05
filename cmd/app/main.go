package main

import (
	"log"

	"Rest_go/internal/database"
	"Rest_go/internal/handlers"
	"Rest_go/internal/models"

	"Rest_go/internal/tasksService"
	"Rest_go/internal/userService"
	"Rest_go/internal/web/tasks"
	"Rest_go/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Проверка ошибки при автоматической миграции
	if err := database.DB.AutoMigrate(&models.Task{}, &models.User{}); err != nil {

		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Создание репозиториев, сервисов и обработчиков
	tasksRepo := tasksService.NewTaskRepository(database.DB)
	tasksService := tasksService.NewTasksService(tasksRepo) // исправлено название метода
	tasksHandler := handlers.NewTaskHandler(tasksService)   // исправлено название функции

	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo) // исправлено название метода
	userHandler := handlers.NewUserHandler(userService) // исправлено название функции

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

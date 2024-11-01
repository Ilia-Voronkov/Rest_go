package main

import (
	"log"

	"Rest_go/internal/database"
	"Rest_go/internal/handlers"
	"Rest_go/internal/tasksService"
	"Rest_go/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Проверяем результат AutoMigrate на наличие ошибок
	if err := database.DB.AutoMigrate(&tasksService.Task{}); err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	// Создание репозитория, сервиса и хендлеров
	repo := tasksService.NewTaskRepository(database.DB)
	service := tasksService.NewService(repo)
	handler := handlers.NewHandler(service)

	// Инициализация Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрация маршрутов
	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

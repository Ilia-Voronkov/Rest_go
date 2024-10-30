package main

import (
	"Rest_go/internal/database"
	"Rest_go/internal/handlers"
	"Rest_go/internal/messagesService"
	"Rest_go/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	// Создание репозитория, сервиса и обработчика
	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)
	handler := handlers.NewHandler(service)

	// Инициализация Echo
	e := echo.New()

	// Включение middleware для логгирования и восстановления после ошибок
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Подключение сгенерированных обработчиков
	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

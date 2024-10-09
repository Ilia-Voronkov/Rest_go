package main

import (
	"log"
	"net/http"

	"Rest_go/internal/database"
	"Rest_go/internal/handlers"
	"Rest_go/internal/messagesService"
	"github.com/gorilla/mux"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	// Создание репозитория, сервиса и хендлеров
	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)
	handler := handlers.NewHandler(service)

	// Настройка роутера
	router := mux.NewRouter()
	router.HandleFunc("/api/message", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/message", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/message/{id}", handler.PatchMessageHandler).Methods("PATCH")
	router.HandleFunc("/api/message/{id}", handler.PutMessageHandler).Methods("PUT")
	router.HandleFunc("/api/message/{id}", handler.DeleteMessageHandler).Methods("DELETE")

	// Запуск сервера
	log.Fatal(http.ListenAndServe(":8080", router))
}

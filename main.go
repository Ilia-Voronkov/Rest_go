package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// requestBody - структура для парсинга JSON запроса
type requestBody struct {
	Message string `json:"message"`
}

// HelloHandler - GET ручка для вывода всех сообщений из БД
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	// Получаем все сообщения из БД
	DB.Find(&messages)

	// Преобразуем массив сообщений в JSON и отправляем в ответе
	err := json.NewEncoder(w).Encode(messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// PostHandler - POST ручка для записи сообщения в БД
func PostHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Создаем новое сообщение и сохраняем его в базу данных
	newMessage := Message{Text: body.Message}

	// Проверяем, возникла ли ошибка при добавлении в базу данных
	err = DB.Create(&newMessage).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error inserting message into DB: %v", err)
		return
	}

	// Возвращаем подтверждение, что сообщение сохранено
	fmt.Fprintf(w, "Message received and saved: %s", body.Message)
}

func main() {
	InitDB()                   // Инициализация БД
	DB.AutoMigrate(&Message{}) // Автоматическая миграция структуры Message

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")   // GET ручка для получения всех сообщений
	router.HandleFunc("/api/message", PostHandler).Methods("POST") // POST ручка для записи нового сообщения

	log.Fatal(http.ListenAndServe(":8080", router)) // Запускаем сервер на порту 8080
}

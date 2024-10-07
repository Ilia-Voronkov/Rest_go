package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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
	err = DB.Create(&newMessage).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error inserting message into DB: %v", err)
		return
	}

	// Возвращаем подтверждение, что сообщение сохранено
	fmt.Fprintf(w, "Message received and saved: %s", body.Message)
}

// PatchHandler - PATCH ручка для обновления сообщения по ID
func PatchHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	var body requestBody
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var message Message
	err = DB.First(&message, id).Error
	if err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	// Обновляем текст сообщения
	message.Text = body.Message
	err = DB.Save(&message).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Message with ID %d updated", id)
}

// PutHandler - PUT ручка для полного обновления сообщения по ID
func PutHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	var body requestBody
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var message Message
	err = DB.First(&message, id).Error
	if err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	// Полностью заменяем текст сообщения
	message.Text = body.Message
	err = DB.Save(&message).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Message with ID %d fully updated", id)
}

// DeleteHandler - DELETE ручка для удаления сообщения по ID
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	var message Message
	err = DB.First(&message, id).Error
	if err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	err = DB.Delete(&message).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Message with ID %d deleted", id)
}

func main() {
	InitDB()                   // Инициализация БД
	DB.AutoMigrate(&Message{}) // Автоматическая миграция структуры Message

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")            // GET ручка для получения всех сообщений
	router.HandleFunc("/api/message", PostHandler).Methods("POST")          // POST ручка для записи нового сообщения
	router.HandleFunc("/api/message/{id}", PatchHandler).Methods("PATCH")   // PATCH ручка для частичного обновления сообщения
	router.HandleFunc("/api/message/{id}", PutHandler).Methods("PUT")       // PUT ручка для полного обновления сообщения
	router.HandleFunc("/api/message/{id}", DeleteHandler).Methods("DELETE") // DELETE ручка для удаления сообщения

	log.Fatal(http.ListenAndServe(":8080", router)) // Запускаем сервер на порту 8080
}

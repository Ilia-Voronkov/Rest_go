package messagesService

import "gorm.io/gorm"

// Message - структура для хранения сообщений в БД
type Message struct {
	gorm.Model
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"` // Добавлено поле IsDone
}

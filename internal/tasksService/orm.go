package tasksService

import (
	"Rest_go/internal/models" // Импортируем новый пакет models
	"gorm.io/gorm"
)

// Task - структура задачи
type Task struct {
	gorm.Model
	Text   string      `json:"text"`
	IsDone bool        `json:"is_done"`
	UserID uint        `json:"user_id"` // ID пользователя, владеющего задачей
	User   models.User `json:"user"`    // Связь с пользователем
}

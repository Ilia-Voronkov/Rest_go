package tasksService

import "gorm.io/gorm"

// Task - структура для хранения задач в БД
type Task struct {
	gorm.Model
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"` // Поле для обозначения завершенности задачи
}

package models

import "gorm.io/gorm"

// User - структура пользователя
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Tasks    []Task `json:"tasks"` // Связь с задачами
}

// Task - структура задачи
type Task struct {
	gorm.Model
	Text   string `json:"text"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"` // ID пользователя, владеющего задачей
}

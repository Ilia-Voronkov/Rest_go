package userService

import "Rest_go/internal/models" // Импортируем новый пакет models

// User - структура пользователя
type User struct {
	ID       uint          `json:"id" gorm:"primaryKey"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Tasks    []models.Task `json:"tasks"` // Связь с задачами

}

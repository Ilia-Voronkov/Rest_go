package userService

import "gorm.io/gorm"

// User - структура для хранения пользователей в БД
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

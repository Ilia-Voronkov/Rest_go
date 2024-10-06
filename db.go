package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Переменная, через которую мы будем работать с БД
var DB *gorm.DB

// InitDB - инициализация подключения к базе данных
func InitDB() {
	// Подключаемся к базе данных, используя параметры, указанные при создании контейнера
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}

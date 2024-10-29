# Переменные для подключения к базе данных
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Команда для создания новой миграции без указания версии или последовательности
migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

# Применение всех миграций
migrate:
	$(MIGRATE) up

# Откат всех миграций
migrate-down:
	$(MIGRATE) down

# Запуск сервера
run:
	go run cmd/app/main.go

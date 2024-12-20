# Переменные для подключения к базе данных
DB_DSN := "postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Команда для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations -seq $(NAME)

# Применение всех миграций
migrate:
	$(MIGRATE) up

# Откат всех миграций
migrate-down:
	$(MIGRATE) down

# Запуск сервера
run:
	go run cmd/app/main.go


lint:
	golangci-lint run --out-format=colored-line-number
gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

package userService

import (
	"Rest_go/internal/models"
)

// UserService - структура для бизнес-логики пользователей
type UserService struct {
	repo UserRepository
}

// NewUserService - конструктор UserService
func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetAllUsers - получение всех пользователей
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

// CreateUser - создание нового пользователя
func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.repo.CreateUser(user)
}

// UpdateUser - обновление данных пользователя
func (s *UserService) UpdateUser(id uint, updatedUser models.User) (models.User, error) {
	return s.repo.UpdateUser(id, updatedUser)
}

// DeleteUserByID - удаление пользователя по ID
func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}

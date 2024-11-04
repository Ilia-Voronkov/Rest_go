package userService

import (
	"Rest_go/internal/models"
	"gorm.io/gorm"
)

// UserRepository - интерфейс для доступа к пользователям в базе данных
type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(id uint, updatedUser models.User) (models.User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository - конструктор userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Реализация метода GetAllUsers
func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Реализация метода CreateUser
func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Реализация метода UpdateUser
func (r *userRepository) UpdateUser(id uint, updatedUser models.User) (models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return models.User{}, err
	}
	if err := r.db.Model(&user).Updates(updatedUser).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Реализация метода DeleteUserByID
func (r *userRepository) DeleteUserByID(id uint) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

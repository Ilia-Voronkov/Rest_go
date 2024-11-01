package userService

import "gorm.io/gorm"

// UserRepository - интерфейс для доступа к пользователям в базе данных
type UserRepository interface {
	GetAllUsers() ([]User, error)
	CreateUser(user User) (User, error)
	UpdateUser(id uint, updatedUser User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) UpdateUser(id uint, updatedUser User) (User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return User{}, err
	}
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password
	err = r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) DeleteUserByID(id uint) error {
	return r.db.Delete(&User{}, id).Error
}

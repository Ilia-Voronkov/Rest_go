package tasksService

import (
	"Rest_go/internal/models"
	"gorm.io/gorm"
)

// TaskRepository - интерфейс для доступа к задачам в базе данных
type TaskRepository interface {
	GetAllTasks() ([]models.Task, error)
	CreateTask(task models.Task) (models.Task, error)
	UpdateTask(id uint, updatedTask models.Task) (models.Task, error)
	DeleteTaskByID(id uint) error
	GetTasksByUserID(userID uint) ([]models.Task, error)
}

type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository - конструктор taskRepository
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// Реализация метода GetAllTasks
func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// Реализация метода CreateTask
func (r *taskRepository) CreateTask(task models.Task) (models.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

// Реализация метода UpdateTask
func (r *taskRepository) UpdateTask(id uint, updatedTask models.Task) (models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return models.Task{}, err
	}
	if err := r.db.Model(&task).Updates(updatedTask).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

// Реализация метода DeleteTaskByID
func (r *taskRepository) DeleteTaskByID(id uint) error {
	if err := r.db.Delete(&models.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Реализация метода GetTasksByUserID
func (r *taskRepository) GetTasksByUserID(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil

}

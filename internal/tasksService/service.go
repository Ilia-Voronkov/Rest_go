package tasksService

import "Rest_go/internal/models"

// TasksService - структура для бизнес-логики задач
type TasksService struct {
	repo TaskRepository
}

// NewTasksService - конструктор TasksService
func NewTasksService(repo TaskRepository) *TasksService {
	return &TasksService{repo: repo}
}

// GetAllTasks - получение всех задач
func (s *TasksService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks() // Убедитесь, что метод возвращает []models.Task
}

// CreateTask - создание новой задачи с привязкой к пользователю
func (s *TasksService) CreateTask(task models.Task, userID uint) (models.Task, error) {
	task.UserID = userID
	return s.repo.CreateTask(task) // Приведите типы в соответствие
}

// GetAllTasksForUser - получение всех задач для конкретного пользователя
func (s *TasksService) GetAllTasksForUser(userID uint) ([]models.Task, error) {
	return s.repo.GetTasksByUserID(userID) // Проверьте, что метод возвращает []models.Task
}

// UpdateTask - обновление задачи по ID
func (s *TasksService) UpdateTask(id uint, updatedTask models.Task) (models.Task, error) {
	return s.repo.UpdateTask(id, updatedTask) // Приведите типы в соответствие
}

// DeleteTaskByID - удаление задачи по ID

func (s *TasksService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

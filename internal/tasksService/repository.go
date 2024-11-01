package tasksService

import "gorm.io/gorm"

// TaskRepository - интерфейс для доступа к задачам в базе данных
type TaskRepository interface {
	GetAllTasks() ([]Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTask(id uint, updatedTask Task) (Task, error)
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *taskRepository) UpdateTask(id uint, updatedTask Task) (Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return Task{}, err
	}
	task.Text = updatedTask.Text
	task.IsDone = updatedTask.IsDone
	err = r.db.Save(&task).Error
	return task, err
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	return r.db.Delete(&Task{}, id).Error
}

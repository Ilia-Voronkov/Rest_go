package tasksService

import "gorm.io/gorm"

// TaskRepository - интерфейс репозитория для задач
type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTask(id int, updatedTask Task) (Task, error)
	DeleteTaskByID(id int) error
}

type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository - создает новый репозиторий задач
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTask(id int, updatedTask Task) (Task, error) {
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

func (r *taskRepository) DeleteTaskByID(id int) error {
	return r.db.Delete(&Task{}, id).Error
}

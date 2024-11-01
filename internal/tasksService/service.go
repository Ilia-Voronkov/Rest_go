package tasksService

// TasksService - структура для бизнес-логики задач
type TasksService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TasksService {
	return &TasksService{repo: repo}
}

func (s *TasksService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TasksService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TasksService) UpdateTask(id uint, updatedTask Task) (Task, error) {
	return s.repo.UpdateTask(id, updatedTask)
}

func (s *TasksService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

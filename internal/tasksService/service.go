package tasksService

// TaskService - сервис для работы с задачами
type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTask(id int, newTask Task) (Task, error) {
	return s.repo.UpdateTask(id, newTask)
}

func (s *TaskService) DeleteTaskByID(id int) error {
	return s.repo.DeleteTaskByID(id)
}

package handlers

import (
	"Rest_go/internal/tasksService"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TasksHandler struct {
	Service *tasksService.TasksService
}

func NewTasksHandler(service *tasksService.TasksService) *TasksHandler {
	return &TasksHandler{Service: service}
}

// Метод для получения всех задач
func (h *TasksHandler) GetTasks(ctx echo.Context) error {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, tasks)
}

// Метод для создания новой задачи
func (h *TasksHandler) PostTasks(ctx echo.Context) error {
	var task tasksService.Task
	if err := ctx.Bind(&task); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, createdTask)
}

// Метод для обновления задачи по ID
func (h *TasksHandler) PatchTasksId(ctx echo.Context, id int) error {
	var updatedTask tasksService.Task
	if err := ctx.Bind(&updatedTask); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	updated, err := h.Service.UpdateTask(uint(id), updatedTask)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, updated)
}

// Метод для удаления задачи по ID
func (h *TasksHandler) DeleteTasksId(ctx echo.Context, id int) error {
	err := h.Service.DeleteTaskByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}

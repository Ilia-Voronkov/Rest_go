package handlers

import (
	"Rest_go/internal/models"
	"Rest_go/internal/tasksService"
	"github.com/labstack/echo/v4"
	"net/http"
)

// TaskHandler - структура обработчика задач
type TaskHandler struct {
	Service *tasksService.TasksService
}

// NewTaskHandler - конструктор TaskHandler
func NewTaskHandler(service *tasksService.TasksService) *TaskHandler {
	return &TaskHandler{Service: service}
}

// GetTasks - получение всех задач
func (h *TaskHandler) GetTasks(ctx echo.Context) error {

	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, tasks)
}

// PostTasks - создание новой задачи
func (h *TaskHandler) PostTasks(ctx echo.Context) error {
	var task models.Task
	if err := ctx.Bind(&task); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	createdTask, err := h.Service.CreateTask(task, task.UserID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, createdTask)
}

// PatchTasksId - обновление задачи по ID
func (h *TaskHandler) PatchTasksId(ctx echo.Context, id int) error {
	var updatedTask models.Task
	if err := ctx.Bind(&updatedTask); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updatedTask, err := h.Service.UpdateTask(uint(id), updatedTask)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, updatedTask)
}

// DeleteTasksId - удаление задачи по ID (исправлен тип аргумента на int)
func (h *TaskHandler) DeleteTasksId(ctx echo.Context, id int) error {
	if err := h.Service.DeleteTaskByID(uint(id)); err != nil {

		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}

package handlers

import (
	"Rest_go/internal/tasksService"
	"Rest_go/internal/web/tasks"
	"context"
)

type Handler struct {
	Service *tasksService.TaskService
}

func NewHandler(service *tasksService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := tasksService.Task{
		Text:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	updatedData := request.Body
	id := request.Id

	updatedTask, err := h.Service.UpdateTask(id, tasksService.Task{
		Text:   *updatedData.Task,
		IsDone: *updatedData.IsDone,
	})
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.Service.DeleteTaskByID(request.Id)
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

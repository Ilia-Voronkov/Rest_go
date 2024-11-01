package handlers

import (
	"Rest_go/internal/userService"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// Метод для получения всех пользователей
func (h *UserHandler) GetUsers(ctx echo.Context) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, users)
}

// Метод для создания нового пользователя
func (h *UserHandler) PostUsers(ctx echo.Context) error {
	var user userService.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	createdUser, err := h.Service.CreateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, createdUser)
}

// Метод для обновления пользователя по ID
func (h *UserHandler) PatchUsersId(ctx echo.Context, id int) error {
	var updatedUser userService.User
	if err := ctx.Bind(&updatedUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	updated, err := h.Service.UpdateUser(uint(id), updatedUser)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, updated)
}

// Метод для удаления пользователя по ID
func (h *UserHandler) DeleteUsersId(ctx echo.Context, id int) error {
	err := h.Service.DeleteUserByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}

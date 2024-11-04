package handlers

import (
	"Rest_go/internal/models"
	"Rest_go/internal/userService"
	"github.com/labstack/echo/v4"
	"net/http"
)

// UserHandler - структура обработчика пользователей
type UserHandler struct {
	Service *userService.UserService
}

// NewUserHandler - конструктор UserHandler
func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// GetUsers - получение всех пользователей
func (h *UserHandler) GetUsers(ctx echo.Context) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, users)
}

// PostUsers - создание нового пользователя
func (h *UserHandler) PostUsers(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	createdUser, err := h.Service.CreateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, createdUser)
}

// PatchUsersId - обновление данных пользователя по ID
func (h *UserHandler) PatchUsersId(ctx echo.Context, id int) error {
	var updatedUser models.User
	if err := ctx.Bind(&updatedUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	updatedUser, err := h.Service.UpdateUser(uint(id), updatedUser)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, updatedUser)
}

// DeleteUsersId - удаление пользователя по ID (исправлен тип аргумента на int)
func (h *UserHandler) DeleteUsersId(ctx echo.Context, id int) error {
	if err := h.Service.DeleteUserByID(uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.NoContent(http.StatusNoContent)
}

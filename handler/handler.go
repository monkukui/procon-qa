package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/x-color/simple-webapp/model"
)

func AddTodo(c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return err
	}

	if todo.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

    todo.UID = uid
    model.CreateTodo(todo)

	return c.JSON(http.StatusCreated, todo)
}

func GetTodos(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	todos := model.FindTodos(&model.Todo{UID: uid})
	return c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.DeleteTodo(&model.Todo{ID: todoID, UID: uid}); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdateTodo(c echo.Context) error {
    uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

    todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

    todos := model.FindTodos(&model.Todo{ID: todoID, UID: uid})
    if len(todos) == 0 {
        return echo.ErrNotFound
    }
    todo := todos[0]
    todo.Completed = !todos[0].Completed
    if err := model.UpdateTodo(&todo); err != nil {
		return echo.ErrNotFound
	}

    return c.NoContent(http.StatusNoContent)
}

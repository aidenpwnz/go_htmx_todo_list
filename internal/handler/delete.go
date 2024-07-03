package handler

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/aidenpwnz/todo_list_go/internal/db"
	"github.com/aidenpwnz/todo_list_go/internal/views"
)

func (h *Handler) RenderDeleteTodo(c echo.Context) error {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := db.DeleteTodoItem(ctx, h.DBClient, id)
	if err != nil {
		message := "Failed to delete item"
		return Render(c, views.ErrorAlert(message))
	}

	items, err := db.GetTodoItems(h.DBClient)
	if err != nil {
		message := "Failed to retrieve items"
		return Render(c, views.ErrorAlert(message))
	}
	h.Items = items

	renderErr := Render(c, views.TodoList(h.Items))
	if renderErr != nil {
		return err
	}
	return Render(c, views.SuccessAlert("Item deleted successfully"))
}

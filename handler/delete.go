package handler

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"

	"todo_list/db"
	"todo_list/views"
)

func (h *TodoHandler) RenderDeleteTodo(c echo.Context) error {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := db.DeleteTodoItem(ctx, h.DBClient, id)
	if err != nil {
		message := "Failed to delete item"
		Render(c, views.ErrorAlert(message))
		return nil
	}

	items, err := db.GetTodoItems(h.DBClient)
	if err != nil {
		message := "Failed to retrieve items"
		Render(c, views.ErrorAlert(message))
		return nil
	}
	h.Items = &items

	Render(c, views.SuccessAlert("Item deleted successfully"))
	return Render(c, views.TodoList(*h.Items))
}

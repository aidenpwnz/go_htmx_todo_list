package handler

import (
	"github.com/labstack/echo/v4"

	"todo_list/models"
	"todo_list/views/todos"
)

func (h *TodoHandler) RenderAddTodo(c echo.Context) error {
	r := c.Request()
	r.ParseForm()
	title := r.FormValue("title")
	desc := r.FormValue("description")
	if title != "" {
		newItem := models.TodoItem{Id: models.GenerateID(), Title: title, Description: desc}
		*h.Items = append(*h.Items, newItem)

		return render(c, todos.TodoItem(newItem))
	}

	return nil
}

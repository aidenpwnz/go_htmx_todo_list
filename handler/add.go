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

		// // Render the components to a buffer
		// buf := new(bytes.Buffer)
		// todos.TodoList([]models.TodoItem{newItem}).Render(c.Request().Context(), buf)
		// alert.InfoAlert(title).Render(c.Request().Context(), buf)

		// // Stream the buffer to the response
		// return c.Stream(200, "text/html", buf)

		return render(c, todos.TodoItem(newItem))
	}

	return nil
}

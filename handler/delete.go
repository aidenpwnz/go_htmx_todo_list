package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"todo_list/views/todos"
)

func (h *TodoHandler) RenderDeleteTodo(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	for i, item := range *h.Items {
		if item.Id == id {
			*h.Items = append((*h.Items)[:i], (*h.Items)[i+1:]...)
			break
		}
	}

	return render(c, todos.TodoList(*h.Items))
}

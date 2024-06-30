package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/aidenpwnz/todo_list_go/db"
)

func (h *TodoHandler) GetTodoItems(c echo.Context) error {
	items, err := db.GetTodoItems(h.DBClient)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, items)
}

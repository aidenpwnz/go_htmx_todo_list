package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"todo_list/models"
)

func render(c echo.Context, comp templ.Component) error {
	return comp.Render(c.Request().Context(), c.Response())
}

type TodoHandler struct {
	Items *[]models.TodoItem
}

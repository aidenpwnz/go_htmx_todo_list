package handler

import (
	"github.com/labstack/echo/v4"

	"todo_list/views"
)

func (h *TodoHandler) RenderIndex(c echo.Context) error {
	return Render(c, views.Index(*h.Items))
}

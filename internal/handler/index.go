package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/aidenpwnz/todo_list_go/internal/views"
)

func (h *Handler) RenderIndex(c echo.Context) error {
	return Render(c, views.Index(*h.Items))
}

package handler

import (
	"github.com/labstack/echo/v4"

	"todo_list/views/index"
)

func (h *TodoHandler) RenderIndex(c echo.Context) error {
	return render(c, index.Index(*h.Items))
}

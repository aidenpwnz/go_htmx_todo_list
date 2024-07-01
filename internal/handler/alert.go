package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/aidenpwnz/todo_list_go/internal/views"
)

func (h *Handler) RenderAlert(c echo.Context) error {
	c.Request().ParseForm()
	severity := c.FormValue("alert-severity")
	message := c.FormValue("alert-message")

	fmt.Printf("Received alert request. Severity: %s, Message: %s\n", severity, message)

	switch severity {
	case "success":
		return Render(c, views.SuccessAlert(message))
	case "error":
		return Render(c, views.ErrorAlert(message))
	case "warning":
		return Render(c, views.WarningAlert(message))
	case "info":
		return Render(c, views.InfoAlert(message))
	default:
		return nil
	}
}

func (h *Handler) RemoveAlert(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

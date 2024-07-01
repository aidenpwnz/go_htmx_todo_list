package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/aidenpwnz/todo_list_go/internal/models"
)

func Render(c echo.Context, comp templ.Component) error {
	return comp.Render(c.Request().Context(), c.Response())
}

type Handler struct {
	Items    *[]models.TodoItem
	DBClient *mongo.Client
}

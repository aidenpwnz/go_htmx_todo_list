package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"

	"todo_list/models"
)

func Render(c echo.Context, comp templ.Component) error {
	return comp.Render(c.Request().Context(), c.Response())
}

type TodoHandler struct {
	Items    *[]models.TodoItem
	DBClient *mongo.Client
}

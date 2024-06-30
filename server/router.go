package server

import (
	"github.com/labstack/echo/v4"

	"todo_list/handler"
)

type Router struct {
	app         *echo.Echo
	todoHandler *handler.TodoHandler
}

func NewRouter(app *echo.Echo, todoHandler *handler.TodoHandler) *Router {
	return &Router{
		app:         app,
		todoHandler: todoHandler,
	}
}

func (r *Router) RegisterRoutes() {
	// Main page
	r.app.GET("/", r.todoHandler.RenderIndex)

	// Item CRUD
	r.app.POST("/add", r.todoHandler.RenderAddTodo)
	r.app.DELETE("/delete/:id", r.todoHandler.RenderDeleteTodo)

	// Alerts
	r.app.GET("/alert", r.todoHandler.RenderAlert)
	r.app.DELETE("/remove-alert", r.todoHandler.RemoveAlert)
}

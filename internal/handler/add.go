package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/aidenpwnz/todo_list_go/internal/db"
	"github.com/aidenpwnz/todo_list_go/internal/models"
	"github.com/aidenpwnz/todo_list_go/internal/views"
)

func (h *Handler) RenderAddTodo(c echo.Context) error {
	r := c.Request()
	title := r.FormValue("title")
	desc := r.FormValue("description")
	if title != "" {
		newItem := models.TodoItem{Id: models.GenerateID(), Title: title, Description: desc}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := db.InsertTodoItem(ctx, h.DBClient, newItem)
		if err != nil {
			fmt.Println(err)
			message := "Failed to add item"
			return Render(c, views.ErrorAlert(message))
		}
		// *h.Items = append(*h.Items, newItem)
		items, err := db.GetTodoItems(h.DBClient)
		if err != nil {
			message := "Failed to retrieve items"
			return Render(c, views.ErrorAlert(message))
		}
		h.Items = items

		message := "Item added successfully"
		renderErr := Render(c, views.TodoItem(newItem))
		if renderErr != nil {
			return err
		}
		return Render(c, views.SuccessAlert(message))
	} else {
		message := "Title is required"
		return Render(c, views.ErrorAlert(message))
	}
}

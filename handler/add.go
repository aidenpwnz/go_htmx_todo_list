package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/aidenpwnz/todo_list_go/db"
	"github.com/aidenpwnz/todo_list_go/models"
	"github.com/aidenpwnz/todo_list_go/views"
)

func (h *TodoHandler) RenderAddTodo(c echo.Context) error {
	r := c.Request()
	r.ParseForm()
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
			Render(c, views.ErrorAlert(message))
			return nil
		}
		// *h.Items = append(*h.Items, newItem)
		items, err := db.GetTodoItems(h.DBClient)
		if err != nil {
			message := "Failed to retrieve items"
			Render(c, views.ErrorAlert(message))
			return nil
		}
		h.Items = &items

		message := "Item added successfully"
		Render(c, views.SuccessAlert(message))
		return Render(c, views.TodoItem(newItem))
	} else {
		message := "Title is required"
		Render(c, views.ErrorAlert(message))
		return nil
	}
}

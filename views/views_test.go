package views

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"

	"github.com/aidenpwnz/todo_list_go/models"
)

func TestTextArea(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		_ = textArea().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}
	textarea := doc.Find("textarea#description")
	if textarea.Length() == 0 {
		t.Error("expected textarea with id 'description' to be rendered, but it wasn't")
	}
	if name, exists := textarea.Attr("name"); !exists || name != "description" {
		t.Error("textarea should have name 'description'")
	}
	if rows, exists := textarea.Attr("rows"); !exists || rows != "4" {
		t.Error("textarea should have 4 rows")
	}
}

func TestInput(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		_ = input().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}
	input := doc.Find("input[type='text'][name='title']")
	if input.Length() == 0 {
		t.Error("expected input with type 'text' and name 'title' to be rendered, but it wasn't")
	}
	if placeholder, exists := input.Attr("placeholder"); !exists || placeholder != "Enter title..." {
		t.Error("input should have placeholder 'Enter title...'")
	}
}

func TestForm(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		_ = Form().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}
	form := doc.Find("form")
	if form.Length() == 0 {
		t.Error("expected form to be rendered, but it wasn't")
	}
	if action, exists := form.Attr("hx-post"); !exists || action != "/add" {
		t.Error("form should have hx-post='/add'")
	}
	if form.Find("input[type='text'][name='title']").Length() == 0 {
		t.Error("form should contain title input")
	}
	if form.Find("textarea#description").Length() == 0 {
		t.Error("form should contain description textarea")
	}
	if form.Find("button[type='submit']").Length() == 0 {
		t.Error("form should contain submit button")
	}
}

func TestIndex(t *testing.T) {
	items := []models.TodoItem{
		{Id: "1", Title: "Test Todo 1", Description: "Description 1"},
		{Id: "2", Title: "Test Todo 2", Description: "Description 2"},
	}
	r, w := io.Pipe()
	go func() {
		_ = Index(items).Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}
	if doc.Find("div#alert-container").Length() == 0 {
		t.Error("expected div with id 'alert-container' to be rendered")
	}
	if doc.Find("div#todo-list").Length() == 0 {
		t.Error("expected div with id 'todo-list' to be rendered")
	}
	if doc.Find("form").Length() == 0 {
		t.Error("expected form to be rendered")
	}
	if todoItems := doc.Find("div#todo-list > div"); todoItems.Length() != len(items) {
		t.Errorf("expected %d todo items, got %d", len(items), todoItems.Length())
	}
}

func TestPage(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		_ = Page().Render(context.Background(), w)
		_ = w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}
	if doc.Find("html").Length() == 0 {
		t.Error("expected html tag to be rendered")
	}
	if doc.Find("head").Length() == 0 {
		t.Error("expected head tag to be rendered")
	}
	if doc.Find("body").Length() == 0 {
		t.Error("expected body tag to be rendered")
	}
	if doc.Find("link[rel='stylesheet'][href='/dist/output.css']").Length() == 0 {
		t.Error("expected CSS link to be rendered")
	}
	if doc.Find("script[src='https://unpkg.com/htmx.org@2.0.0']").Length() == 0 {
		t.Error("expected htmx script to be rendered")
	}
}

func TestTodoItem(t *testing.T) {
	item := models.TodoItem{Id: "1", Title: "Test Todo", Description: "Test Description"}
	var w strings.Builder
	err := TodoItem(item).Render(context.Background(), &w)
	if err != nil {
		t.Fatalf("failed to render todo item: %v", err)
	}
	output := w.String()
	if !strings.Contains(output, item.Title) {
		t.Error("rendered output should contain the todo item title")
	}
	if !strings.Contains(output, item.Description) {
		t.Error("rendered output should contain the todo item description")
	}
	if !strings.Contains(output, "hx-delete=\"/delete/1\"") {
		t.Error("rendered output should contain the delete button with correct hx-delete attribute")
	}
}

func TestTodoList(t *testing.T) {
	items := []models.TodoItem{
		{Id: "1", Title: "Test Todo 1", Description: "Description 1"},
		{Id: "2", Title: "Test Todo 2", Description: "Description 2"},
	}
	var w strings.Builder
	err := TodoList(items).Render(context.Background(), &w)
	if err != nil {
		t.Fatalf("failed to render todo list: %v", err)
	}
	output := w.String()
	for _, item := range items {
		if !strings.Contains(output, item.Title) {
			t.Errorf("rendered output should contain the todo item title: %s", item.Title)
		}
		if !strings.Contains(output, item.Description) {
			t.Errorf("rendered output should contain the todo item description: %s", item.Description)
		}
	}
}

func TestAlerts(t *testing.T) {
	alertTypes := []struct {
		name     string
		function func(string) templ.Component
	}{
		{"SuccessAlert", SuccessAlert},
		{"InfoAlert", InfoAlert},
		{"WarningAlert", WarningAlert},
		{"ErrorAlert", ErrorAlert},
	}

	for _, at := range alertTypes {
		t.Run(at.name, func(t *testing.T) {
			message := "Test message"
			var w strings.Builder
			err := at.function(message).Render(context.Background(), &w)
			if err != nil {
				t.Fatalf("failed to render %s: %v", at.name, err)
			}
			output := w.String()
			if !strings.Contains(output, message) {
				t.Errorf("%s should contain the message", at.name)
			}
			if !strings.Contains(output, "hx-delete=\"/remove-alert\"") {
				t.Errorf("%s should contain hx-delete attribute", at.name)
			}
		})
	}
}

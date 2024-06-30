package views

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/a-h/templ"

	"todo_list/models"
)

// func TestTextArea(t *testing.T) {
// 	// Pipe the rendered template into goquery.
// 	r, w := io.Pipe()
// 	go func() {
// 		_ = textArea().Render(context.Background(), w) // Adjust the function call as necessary
// 		_ = w.Close()
// 	}()
// 	doc, err := goquery.NewDocumentFromReader(r)
// 	if err != nil {
// 		t.Fatalf("failed to read template: %v", err)
// 	}
// 	// Expect the textarea to be present with the correct id.
// 	if doc.Find(`textarea#description`).Length() == 0 {
// 		t.Error("expected textarea with id 'description' to be rendered, but it wasn't")
// 	}
// }

// func TestInput(t *testing.T) {
// 	// Pipe the rendered template into goquery.
// 	r, w := io.Pipe()
// 	go func() {
// 		_ = input().Render(context.Background(), w) // Adjust the function call as necessary
// 		_ = w.Close()
// 	}()
// 	doc, err := goquery.NewDocumentFromReader(r)
// 	if err != nil {
// 		t.Fatalf("failed to read template: %v", err)
// 	}
// 	// Expect the input to be present with the correct type and name.
// 	if doc.Find(`input[type="text"][name="title"]`).Length() == 0 {
// 		t.Error("expected input with type 'text' and name 'title' to be rendered, but it wasn't")
// 	}
// }

// func TestForm(t *testing.T) {
// 	// Pipe the rendered template into goquery.
// 	r, w := io.Pipe()
// 	go func() {
// 		_ = Form().Render(context.Background(), w) // Adjust the function call as necessary
// 		_ = w.Close()
// 	}()
// 	doc, err := goquery.NewDocumentFromReader(r)
// 	if err != nil {
// 		t.Fatalf("failed to read template: %v", err)
// 	}
// 	// Expect the form to be present with the correct attributes.
// 	formSelector := `form[hx-post="/add"]`
// 	if doc.Find(formSelector).Length() == 0 {
// 		t.Errorf("expected form with action '/submit' and method 'post' to be rendered, but it wasn't")
// 	}
// 	// Optionally, check for the presence of child elements like input and textarea.
// 	if doc.Find(formSelector+` input[type="text"][name="title"]`).Length() == 0 {
// 		t.Error("expected input with type 'text' and name 'title' inside the form, but it wasn't")
// 	}
// 	if doc.Find(formSelector+` textarea#description`).Length() == 0 {
// 		t.Error("expected textarea with id 'description' inside the form, but it wasn't")
// 	}
// }

// func TestIndex(t *testing.T) {
// 	// Mock data for TodoList items
// 	items := []models.TodoItem{
// 		{Title: "Test Todo 1", Description: "Description 1"},
// 		// Add more mock items as needed
// 	}

// 	// Pipe the rendered template into goquery.
// 	r, w := io.Pipe()
// 	go func() {
// 		_ = Index(items).Render(context.Background(), w)
// 		_ = w.Close()
// 	}()
// 	doc, err := goquery.NewDocumentFromReader(r)
// 	if err != nil {
// 		t.Fatalf("failed to read template: %v", err)
// 	}

// 	// Verify alert-container div is present
// 	if doc.Find(`div#alert-container`).Length() == 0 {
// 		t.Error("expected div with id 'alert-container' to be rendered, but it wasn't")
// 	}

// 	// Verify the h1 tag with specific classes is present
// 	if doc.Find(`h1.text-2xl.font-extrabold.text-gray-900.dark\:text-white.pt-8.h-\[15vh\]`).Length() == 0 {
// 		t.Error("expected h1 with specific classes to be rendered, but it wasn't")
// 	}

// 	// Verify todo-list div is present
// 	if doc.Find(`div#todo-list`).Length() == 0 {
// 		t.Error("expected div with id 'todo-list' to be rendered, but it wasn't")
// 	}

// 	// Verify Form component is rendered
// 	if doc.Find(`div.container.mx-auto.flex.flex-col.justify-end.items-end.my-4.gap-y-4.h-\[35vh\]`).Find("form").Length() == 0 {
// 		t.Error("expected Form component to be rendered inside the specified div, but it wasn't")
// 	}
// }

// func TestTodoItem(t *testing.T) {
// 	item := models.TodoItem{
// 		Id:          "test-id",
// 		Title:       "Test Title",
// 		Description: "Test Description",
// 	}

// 	var w strings.Builder
// 	err := TodoItem(item).Render(context.Background(), &w) // Assuming a function that renders TodoItem to a string
// 	if err != nil {
// 		t.Fatalf("failed to render todo item: %v", err)
// 	}
// 	output := w.String()

// 	// Verify title and description
// 	if !strings.Contains(output, item.Title) {
// 		t.Errorf("expected title %q to be in output", item.Title)
// 	}
// 	if !strings.Contains(output, item.Description) {
// 		t.Errorf("expected description %q to be in output", item.Description)
// 	}

// 	// Verify delete button
// 	expectedDeleteURL := fmt.Sprintf("/delete/%s", item.Id)
// 	if !strings.Contains(output, expectedDeleteURL) {
// 		t.Errorf("expected delete button with URL %q", expectedDeleteURL)
// 	}
// }

// func TestTodoList(t *testing.T) {
// 	items := []models.TodoItem{
// 		{Id: "id1", Title: "Title 1", Description: "Description 1"},
// 		{Id: "id2", Title: "Title 2", Description: "Description 2"},
// 	}

// 	var w strings.Builder
// 	err := TodoList(items).Render(context.Background(), &w) // Assuming a function that renders TodoList to a string
// 	if err != nil {
// 		t.Fatalf("failed to render todo list: %v", err)
// 	}
// 	output := w.String()

// 	// Verify each TodoItem is rendered
// 	for _, item := range items {
// 		if !strings.Contains(output, item.Title) || !strings.Contains(output, item.Description) {
// 			t.Errorf("expected item with title %q and description %q to be in output", item.Title, item.Description)
// 		}
// 	}
// }

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

package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/aidenpwnz/todo_list_go/internal/handler"
	"github.com/aidenpwnz/todo_list_go/internal/models"
)

func TestSetupServer(t *testing.T) {
	os.Setenv("MONGO_URI", "mongodb://localhost:27017/test")
	defer os.Unsetenv("MONGO_URI")

	app, dbClient := SetupServer()
	defer dbClient.Disconnect(context.Background())

	if app == nil {
		t.Error("Expected app to be initialized, got nil")
	}

	if dbClient == nil {
		t.Error("Expected dbClient to be initialized, got nil")
	}
}

func TestRenderIndex(t *testing.T) {
	os.Setenv("MONGO_URI", "mongodb://localhost:27017/test")
	defer os.Unsetenv("MONGO_URI")

	app, dbClient := SetupServer()
	defer dbClient.Disconnect(context.Background())

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := app.NewContext(req, rec)
	todoHandler := handler.Handler{
		Items:    &[]models.TodoItem{},
		DBClient: dbClient,
	}

	if err := todoHandler.RenderIndex(c); err != nil {
		t.Errorf("RenderIndex() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestRenderAddTodo(t *testing.T) {
	os.Setenv("MONGO_URI", "mongodb://localhost:27017/test")
	defer os.Unsetenv("MONGO_URI")

	app, dbClient := SetupServer()
	defer dbClient.Disconnect(context.Background())

	body := strings.NewReader("title=Test Todo")
	req := httptest.NewRequest(http.MethodPost, "/add", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	c := app.NewContext(req, rec)
	todoHandler := handler.Handler{
		Items:    &[]models.TodoItem{},
		DBClient: dbClient,
	}

	if err := todoHandler.RenderAddTodo(c); err != nil {
		t.Errorf("RenderAddTodo() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestRenderDeleteTodo(t *testing.T) {
	os.Setenv("MONGO_URI", "mongodb://localhost:27017/test")
	defer os.Unsetenv("MONGO_URI")

	app, dbClient := SetupServer()
	defer dbClient.Disconnect(context.Background())

	req := httptest.NewRequest(http.MethodDelete, "/delete/123", nil)
	rec := httptest.NewRecorder()

	c := app.NewContext(req, rec)
	todoHandler := handler.Handler{
		Items:    &[]models.TodoItem{},
		DBClient: dbClient,
	}

	if err := todoHandler.RenderDeleteTodo(c); err != nil {
		t.Errorf("RenderDeleteTodo() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestRenderAlert(t *testing.T) {
	os.Setenv("MONGO_URI", "mongodb://localhost:27017/test")
	defer os.Unsetenv("MONGO_URI")

	app, dbClient := SetupServer()
	defer dbClient.Disconnect(context.Background())

	req := httptest.NewRequest(http.MethodGet, "/alert", nil)
	rec := httptest.NewRecorder()

	c := app.NewContext(req, rec)
	todoHandler := handler.Handler{
		Items:    &[]models.TodoItem{},
		DBClient: dbClient,
	}

	if err := todoHandler.RenderAlert(c); err != nil {
		t.Errorf("RenderAlert() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestRemoveAlert(t *testing.T) {
	os.Setenv("MONGO_URI", "mongodb://localhost:27017/test")
	defer os.Unsetenv("MONGO_URI")

	app, dbClient := SetupServer()
	defer dbClient.Disconnect(context.Background())

	req := httptest.NewRequest(http.MethodDelete, "/remove-alert", nil)
	rec := httptest.NewRecorder()

	c := app.NewContext(req, rec)
	todoHandler := handler.Handler{
		Items:    &[]models.TodoItem{},
		DBClient: dbClient,
	}

	if err := todoHandler.RemoveAlert(c); err != nil {
		t.Errorf("RemoveAlert() error = %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

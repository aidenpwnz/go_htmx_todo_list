package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/aidenpwnz/todo_list_go/internal/models"
	"github.com/aidenpwnz/todo_list_go/internal/views"
)

func TestRenderAddTodo(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		// Mock responses for InsertTodoItem and GetTodoItems
		mt.AddMockResponses(
			mtest.CreateSuccessResponse(), // For InsertTodoItem
			mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
				{Key: "id", Value: "1"},
				{Key: "title", Value: "Test Todo"},
				{Key: "description", Value: "Test Description"},
			}),
		)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader("title=Test+Todo&description=Test+Description"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{DBClient: mt.Client}
		err := h.RenderAddTodo(c)

		assert.NoError(t, err)
		assert.Contains(t, rec.Body.String(), "Test Todo")
		assert.Contains(t, rec.Body.String(), "Test Description")
		assert.Contains(t, rec.Body.String(), "Item added successfully")
	})

	mt.Run("empty title", func(mt *mtest.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader("title=&description=Test+Description"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{DBClient: mt.Client}
		err := h.RenderAddTodo(c)

		assert.NoError(t, err)
		assert.Contains(t, rec.Body.String(), "Title is required")
	})

	mt.Run("database insert error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    12345,
			Message: "Database error",
		}))

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader("title=Test+Todo&description=Test+Description"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{DBClient: mt.Client}
		err := h.RenderAddTodo(c)

		assert.NoError(t, err)
		assert.Contains(t, rec.Body.String(), "Failed to add item")
	})

	mt.Run("database get error", func(mt *mtest.T) {
		mt.AddMockResponses(
			mtest.CreateSuccessResponse(), // For InsertTodoItem
			mtest.CreateCommandErrorResponse(mtest.CommandError{
				Code:    12345,
				Message: "Database error",
			}),
		)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader("title=Test+Todo&description=Test+Description"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{DBClient: mt.Client}
		err := h.RenderAddTodo(c)

		assert.NoError(t, err)
		assert.Contains(t, rec.Body.String(), "Failed to retrieve items")
	})
}

func TestRenderAlert(t *testing.T) {
	testCases := []struct {
		name     string
		severity string
		message  string
		expected string
	}{
		{"success", "success", "Success message", "Success!"},
		{"error", "error", "Error message", "Error"},
		{"warning", "warning", "Warning message", "Warning"},
		{"info", "info", "Info message", "Info"},
		{"invalid", "invalid", "Invalid message", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/alert", nil)
			q := req.URL.Query()
			q.Add("alert-severity", tc.severity)
			q.Add("alert-message", tc.message)
			req.URL.RawQuery = q.Encode()
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			h := &Handler{}
			err := h.RenderAlert(c)

			assert.NoError(t, err)
			if tc.expected != "" {
				assert.Contains(t, rec.Body.String(), tc.expected)
				assert.Contains(t, rec.Body.String(), tc.message)
			} else {
				assert.Empty(t, rec.Body.String())
			}
		})
	}
}

func TestRemoveAlert(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/remove-alert", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := &Handler{}
	err := h.RemoveAlert(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Empty(t, rec.Body.String())
}

func TestRenderDeleteTodo(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse(), mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch))

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/delete/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/delete/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		h := &Handler{DBClient: mt.Client}
		err := h.RenderDeleteTodo(c)

		assert.NoError(t, err)
		assert.Contains(t, rec.Body.String(), "Item deleted successfully")
	})

	mt.Run("delete error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    12345,
			Message: "Delete error",
		}))

		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/delete/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/delete/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		h := &Handler{DBClient: mt.Client}
		err := h.RenderDeleteTodo(c)

		assert.NoError(t, err)
		assert.Contains(t, rec.Body.String(), "Failed to delete item")
	})
}

func TestRenderIndex(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	items := []models.TodoItem{
		{Id: "1", Title: "Test Todo", Description: "Test Description"},
	}
	h := &Handler{Items: &items}
	err := h.RenderIndex(c)

	assert.NoError(t, err)
	assert.Contains(t, rec.Body.String(), "Test Todo")
	assert.Contains(t, rec.Body.String(), "Test Description")
}

func TestGetTodoItems(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		expectedItems := []models.TodoItem{
			{Id: "1", Title: "Test Todo", Description: "Test Description"},
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "id", Value: expectedItems[0].Id},
			{Key: "title", Value: expectedItems[0].Title},
			{Key: "description", Value: expectedItems[0].Description},
		}))

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/items", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{DBClient: mt.Client}
		err := h.GetTodoItems(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var items []models.TodoItem
		err = json.Unmarshal(rec.Body.Bytes(), &items)
		assert.NoError(t, err)
		assert.Equal(t, expectedItems, items)
	})

	mt.Run("database error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    12345,
			Message: "Database error",
		}))

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/items", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		h := &Handler{DBClient: mt.Client}
		err := h.GetTodoItems(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "Database error")
	})
}

func TestRender(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	testComponent := views.SuccessAlert("Test message")
	err := Render(c, testComponent)

	assert.NoError(t, err)
	assert.Contains(t, rec.Body.String(), "Test message")
	assert.Contains(t, rec.Body.String(), "Success!")
}

package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	"github.com/aidenpwnz/todo_list_go/internal/models"
)

func TestConnect(t *testing.T) {
	t.Run("ValidURI", func(t *testing.T) {
		uri := "mongodb://localhost:27017"
		client, err := Connect(uri)
		assert.NoError(t, err)
		assert.NotNil(t, client)
		defer client.Disconnect(context.Background())
	})

	t.Run("InvalidURI", func(t *testing.T) {
		uri := "invalid://localhost:27017"
		client, err := Connect(uri)
		assert.Error(t, err)
		assert.Nil(t, client)
	})

	t.Run("Timeout", func(t *testing.T) {
		_, err := Connect("mongodb://invalid-host:27017")
		if err == nil {
			t.Error("An error is expected but got nil.")
		}
	})
}

func TestInsertTodoItem(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		item := models.TodoItem{
			Id:          "1",
			Description: "Test todo item",
			Title:       "Test todo item",
		}
		_, err := mt.Client.Database("mongodb").Collection("todos").InsertOne(context.Background(), item)
		assert.Nil(t, err)
	})

	mt.Run("duplicate key", func(mt *mtest.T) {
		item := models.TodoItem{
			Id:          "1",
			Description: "Test todo item",
			Title:       "Test todo item",
		}

		// Mock response for the first insert (success)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		// Mock response for the second insert (duplicate key error)
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000, // MongoDB duplicate key error code
			Message: "duplicate key error",
		}))

		// First insert should succeed
		err := InsertTodoItem(context.Background(), mt.Client, item)
		if err != nil {
			mt.Errorf("Expected first insert to succeed, got error: %v", err)
		}

		// Second insert should fail with duplicate key error
		err = InsertTodoItem(context.Background(), mt.Client, item)
		if err == nil {
			mt.Error("Expected error for duplicate key, got nil")
		}
	})
}

func TestGetTodoItems(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		expectedTodos := []models.TodoItem{
			{Id: "1", Title: "Test Todo 1", Description: "Description 1"},
			{Id: "2", Title: "Test Todo 2", Description: "Description 2"},
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "id", Value: expectedTodos[0].Id},
			{Key: "title", Value: expectedTodos[0].Title},
			{Key: "description", Value: expectedTodos[0].Description},
		}, bson.D{
			{Key: "id", Value: expectedTodos[1].Id},
			{Key: "title", Value: expectedTodos[1].Title},
			{Key: "description", Value: expectedTodos[1].Description},
		}))

		todos, err := GetTodoItems(mt.Client)
		assert.Nil(mt, err)
		assert.Equal(mt, expectedTodos, todos)
	})

	mt.Run("error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    1,
			Message: "mock error",
		}))

		_, err := GetTodoItems(mt.Client)

		if err == nil {
			mt.Errorf("Expected error from mock response, got nil")
		}
	})
}

func TestDeleteTodoItem(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := DeleteTodoItem(context.Background(), mt.Client, "1")
		if err != nil {
			mt.Errorf("Expected no error, got: %v", err)
		}
	})

	mt.Run("not found", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "item not found",
		}))
		err := DeleteTodoItem(context.Background(), mt.Client, "1")

		if err == nil {
			mt.Errorf("Expected error for not found item, got nil")
		}
	})
}

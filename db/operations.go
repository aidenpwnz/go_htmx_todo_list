package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"todo_list/models"
)

// Create a new todo item
func InsertTodoItem(ctx context.Context, dbClient *mongo.Client, item models.TodoItem) error {
	collection := dbClient.Database("todo_list_go").Collection("todos")
	_, err := collection.InsertOne(ctx, item)
	return err
}

func GetTodoItems(dbClient *mongo.Client) ([]models.TodoItem, error) {
	var items []models.TodoItem
	collection := dbClient.Database("todo_list_go").Collection("todos")

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var item models.TodoItem
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func DeleteTodoItem(ctx context.Context, dbClient *mongo.Client, id string) error {
	collection := dbClient.Database("todo_list_go").Collection("todos")
	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	return err
}

package server

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/aidenpwnz/todo_list_go/db"
	"github.com/aidenpwnz/todo_list_go/handler"
)

func SetupServer() (*echo.Echo, *mongo.Client) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	connString := os.Getenv("MONGO_URI")
	if connString == "" {
		log.Fatal("MONGO_URI must be set in .env")
	}

	// Connect to MongoDB
	dbClient, err := db.Connect(connString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := echo.New()
	app.Static("/dist", "dist")

	data, _ := db.GetTodoItems(dbClient)

	todoHandler := handler.TodoHandler{
		Items:    &data,
		DBClient: dbClient,
	}
	router := NewRouter(app, &todoHandler)
	router.RegisterRoutes()

	return app, dbClient
}

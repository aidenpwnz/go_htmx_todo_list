package main

import (
	"context"
	"fmt"

	"github.com/aidenpwnz/todo_list_go/internal/server"
)

func main() {
	app, dbClient := server.SetupServer()

	defer func() {
		err := dbClient.Disconnect(context.Background())
		if err != nil {
			fmt.Printf("Failed to disconnect from the database: %v\n", err)
		}
	}()

	port := ":8080"

	app.Logger.Fatal(app.Start(port))

	fmt.Printf("Server Started at port %s\n", port)
}

package main

import (
	"context"
	"fmt"

	"github.com/aidenpwnz/todo_list_go/internal/server"
)

func main() {
	app, dbClient := server.SetupServer()

	defer dbClient.Disconnect(context.Background())

	port := ":8080"

	app.Logger.Fatal(app.Start(port))

	fmt.Printf("Server Started at port %s\n", port)
}

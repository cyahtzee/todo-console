package main

import (
	"fmt"

	"todo-console/app"
)

type Item interface {
	NewTodo() *app.Todo
	GetID() int
	GetTitle() string
	GetDescription() string
	GetCompleted() bool
}

func main() {
	todos := &[]app.Todo{}
	consoleMenu := app.Menu{Running: true, Items: todos}

	if err := consoleMenu.Run(); err != nil {
		fmt.Println("Error:", err)
	}

	defer func() {
		fmt.Println("Exiting...")
	}()
}

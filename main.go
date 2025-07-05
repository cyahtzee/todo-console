package main

import (
	"fmt"
	"todo-console/app"
	"todo-console/service/todo"
	"todo-console/storage"
	"todo-console/types"
)

func main() {
	store := storage.NewStorage(&[]storage.Item{})
	routerCtx := types.NewRouterContext("main", &storage.Todo{})
	router := app.NewRouter(routerCtx)

	// Create concrete tab instances
	mainTab := &todo.MainTab{
		Tab: todo.Tab{
			Active:  true,
			Name:    "main",
			Storage: store,
		},
	}
	listTab := &todo.ListTab{
		Tab: todo.Tab{
			Active:  false,
			Name:    "list",
			Storage: store,
		},
	}
	addTab := &todo.AddTab{
		Tab: todo.Tab{
			Active:  false,
			Name:    "add",
			Storage: store,
		},
	}
	itemTab := &todo.ItemTab{
		Tab: todo.Tab{
			Active:  false,
			Name:    "item",
			Storage: store,
		},
	}

	tabs := &[]todo.TabInterface{mainTab, listTab, addTab, itemTab}

	appInstance := app.NewApp(store, router, tabs)

	if err := appInstance.Run(); err != nil {
		fmt.Println("Error:", err)
	}

	defer func() {
		fmt.Println("Exiting...")
	}()
}

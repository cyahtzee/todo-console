package main

import (
	"fmt"
	"todo-console/app"
	"todo-console/service/todo"
	"todo-console/storage"
)

func main() {
	storage := storage.NewStorage(&[]storage.Item{})
	router := app.NewRouter()

	tabs := &[]todo.TabInterface{
		&todo.MainTab{
			Tab: todo.Tab{
				Active:  true,
				Name:    "main",
				Storage: storage,
			},
		},
		&todo.ListTab{
			Tab: todo.Tab{
				Active:  false,
				Name:    "list",
				Storage: storage,
			},
		},
		&todo.AddTab{
			Tab: todo.Tab{
				Active:  false,
				Name:    "add",
				Storage: storage,
			},
		},
		&todo.ItemTab{
			Tab: todo.Tab{
				Active:  false,
				Name:    "item",
				Storage: storage,
			},
		},
	}

	appInstance := app.NewApp(storage, tabs, router)

	if err := appInstance.Run(); err != nil {
		fmt.Println("Error:", err)
	}

	defer func() {
		fmt.Println("Exiting...")
	}()
}

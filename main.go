package main

import (
	"fmt"

	"todo-console/app"
)

func main() {
	todos := &[]app.Todo{}

	tabs := &[]app.TabInterface{
		&app.MainTab{
			Tab: app.Tab{
				Active: true,
				Name:   "main",
				Items:  todos,
			},
		},
		&app.ListTab{
			Tab: app.Tab{
				Active: false,
				Name:   "list",
				Items:  todos,
			},
		},
		&app.AddTab{
			Tab: app.Tab{
				Active: false,
				Name:   "add",
				Items:  todos,
			},
		},
		&app.ItemTab{
			Tab: app.Tab{
				Active: false,
				Name:   "item",
				Items:  todos,
			},
		},
	}

	appInstance := app.NewApp(todos, tabs)

	if err := appInstance.Run(); err != nil {
		fmt.Println("Error:", err)
	}

	defer func() {
		fmt.Println("Exiting...")
	}()
}

package main

import (
	"fmt"
	"todo-console/app"
	"todo-console/constants"
	"todo-console/service/todo"
	"todo-console/storage"
	"todo-console/types"
)

func main() {
	store := storage.NewCSVStorage(constants.DEFAULT_CSV_FILE_PATH)
	routerCtx := types.NewRouterContext("main", storage.NewTodo(0, "", ""))
	router := app.NewRouter(routerCtx)

	// Create concrete tab instances
	mainTab := &todo.MainTab{
		Tab: todo.Tab{
			Active:  true,
			Name:    constants.TabMain,
			Storage: store,
		},
	}
	listTab := &todo.ListTab{
		Tab: todo.Tab{
			Active:  false,
			Name:    constants.TabList,
			Storage: store,
		},
	}
	addTab := &todo.AddTab{
		Tab: todo.Tab{
			Active:  false,
			Name:    constants.TabAdd,
			Storage: store,
		},
	}
	itemTab := &todo.ItemTab{
		Tab: todo.Tab{
			Active:  false,
			Name:    constants.TabItem,
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

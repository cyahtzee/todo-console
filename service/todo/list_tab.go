package todo

import (
	"fmt"
	"strconv"
	"todo-console/storage"
)

type ListTab struct {
	Tab
}

func (t *ListTab) Open() error {
	t.Tab.Open()
	showAllTodos(t.Storage.FindAll())
	return nil
}

func (t *ListTab) HandleInput(input string) TabInput {
	tabName := ""
	ctx := t.Ctx
	v, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid option. Please try again.")
		return NewTabInput("list", t.Ctx)
	}

	if len(*t.Storage.FindAll()) > 0 && v > 0 {
		ctx = t.Storage.Find(v)
	}

	switch v {
	case 0:
		tabName = "main"
	default:
		tabName = "item"
	}

	return NewTabInput(tabName, ctx)
}

func showAllTodos(items *[]storage.Item) error {
	if len(*items) == 0 {
		fmt.Println("No todos found")
		return nil
	}

	fmt.Println("Your todos:")
	for _, item := range *items {
		fmt.Printf("%d. %s\n", item.GetID(), item.GetTitle())
	}

	return nil
}

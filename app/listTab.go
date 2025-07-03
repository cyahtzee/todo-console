package app

import (
	"fmt"
)

type ListTab struct {
	Tab
}

func (t *ListTab) Open() error {
	t.Tab.Open()
	showAllTodos(t.Items)
	return nil
}

func (t *ListTab) HandleInput(input string) (string, string) {
	tabName := ""
	ctx := ""

	switch input {
	case "0":
		tabName = "main"
	default:
		tabName = "item"
		ctx = input
	}

	return tabName, ctx
}

func showAllTodos(items *[]Todo) error {
	if len(*items) == 0 {
		fmt.Println("No todos found")
		return nil
	}

	fmt.Println("Your todos:")
	for i, item := range *items {
		fmt.Printf("%d. %s\n", i+1, item.GetTitle())
	}

	return nil
}

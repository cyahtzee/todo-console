package app

import (
	"fmt"
)

type AddTab struct {
	Tab
}

func (t *AddTab) HandleInput(input string) (string, string) {
	tabName, err := t.addNewTodo(input)
	ctx := ""

	if err != nil {
		return "add", ctx
	}

	return tabName, ctx
}

func (t *AddTab) addNewTodo(input string) (string, error) {
	if input == "" {
		fmt.Println("Todo is empty")
		return "", fmt.Errorf("todo is empty")
	}

	if input == "0" {
		return "main", nil
	}

	*t.Items = append(*t.Items, *NewTodo(len(*t.Items)+1, input, "", false))
	fmt.Printf("Added: %s\n", input)
	return "list", nil
}

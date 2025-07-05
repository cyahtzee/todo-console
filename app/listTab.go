package app

import (
	"fmt"
	"strconv"
)

type ListTab struct {
	Tab
}

func (t *ListTab) Open() error {
	t.Tab.Open()
	showAllTodos(t.Items)
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

	if len(*t.Items) > 0 && v > 0 {
		ctx = &(*t.Items)[v-1]
	}

	switch v {
	case 0:
		tabName = "main"
	default:
		tabName = "item"
	}

	return NewTabInput(tabName, ctx)
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

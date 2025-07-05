package app

import (
	"fmt"
)

type AddTab struct {
	Tab
}

func (t *AddTab) Open() error {
	t.Tab.Open()
	fmt.Println("Enter the " + t.GetProperty() + ": ")
	return nil
}

func (t *AddTab) HandleInput(input string) TabInput {
	err := t.addNewTodo(input)
	tabName := "list"
	ctx := t.Ctx

	if err != nil || !t.Ctx.Valid() {
		tabName = "add"
	}

	if input == "0" {
		tabName = "main"
	}

	return NewTabInput(tabName, ctx)
}

func (t *AddTab) addNewTodo(input string) error {
	if input == "" {
		fmt.Println("Todo is empty")
		return fmt.Errorf("todo is empty")
	}

	field := t.GetProperty()
	t.Ctx.SetField(field, input)

	if t.Ctx.Valid() {
		*t.Items = append(*t.Items, *t.Ctx)
		t.Step = 0
		fmt.Println("Todo added successfully")
		return nil
	}

	t.NextStep()
	return nil
}

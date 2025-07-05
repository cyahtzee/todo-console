package todo

import (
	"fmt"
)

type AddTab struct {
	Tab
}

var requiredFields = []string{
	"title",
	"description",
	"completed",
}

func (t *AddTab) Open() error {
	t.Tab.Open()
	fmt.Println("Enter the " + t.GetProperty() + ": ")
	return nil
}

func (t *AddTab) HandleInput(input string) TabInput {
	err := t.addNewTodo(input)
	tabName := "list"
	ctx := *t.Ctx

	if err != nil || !ctx.Validate(requiredFields) {
		tabName = "add"
	}

	if input == "0" {
		tabName = "main"
	}

	return NewTabInput(tabName, &ctx)
}

func (t *AddTab) addNewTodo(input string) error {
	if input == "" {
		fmt.Println("Todo is empty")
		return fmt.Errorf("todo is empty")
	}

	field := t.GetProperty()
	currentItem := *t.Ctx
	currentItem.SetField(field, input)

	if currentItem.Validate(requiredFields) {
		t.Storage.Create(currentItem)
		t.Step = 0
		fmt.Println("Todo added successfully")
		return nil
	}

	t.NextStep()
	return nil
}

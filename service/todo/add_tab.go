package todo

import (
	"fmt"
	"todo-console/constants"
	"todo-console/storage"
	"todo-console/types"
)

type AddTab struct {
	Tab
}

var requiredFields = []string{
	"title",
	"description",
	"completed",
}

func (t *AddTab) Open(c *types.RouterContext) error {
	t.Tab.Open(c)
	fmt.Println("Enter the " + t.GetProperty() + ": ")
	return nil
}

func (t *AddTab) HandleInput(c *types.RouterContext) *types.RouterContext {
	input := c.Input
	err := t.addNewTodo(c)

	if c.Item.Validate(requiredFields) {
		t.Step = 0
		c.TabName = constants.TabList
		c.Item = storage.NewTodo(0, "", "")
		return c
	}

	if err != nil {
		c.TabName = constants.TabAdd
	}

	if input == "0" {
		c.TabName = constants.TabMain
	}

	return c
}

func (t *AddTab) addNewTodo(c *types.RouterContext) error {
	if c.Input == "" {
		fmt.Println("Todo is empty")
		return fmt.Errorf("todo is empty")
	}

	field := t.GetProperty()
	currentItem := c.Item
	err := currentItem.SetField(field, c.Input)

	if err != nil {
		fmt.Println("Error setting field: ", err)
		return err
	}

	if currentItem.Validate(requiredFields) {
		t.Storage.Create(currentItem)
		fmt.Println("Todo added successfully")
		return nil
	}

	t.NextStep()
	return nil
}

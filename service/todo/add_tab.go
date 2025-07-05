package todo

import (
	"fmt"
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
	c.TabName = "list"

	if err != nil || !c.Item.Validate(requiredFields) {
		c.TabName = "add"
	}

	if input == "0" {
		c.TabName = "main"
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
	currentItem.SetField(field, c.Input)

	if currentItem.Validate(requiredFields) {
		t.Storage.Create(currentItem)
		t.Step = 0
		fmt.Println("Todo added successfully")
		return nil
	}

	t.NextStep()
	return nil
}

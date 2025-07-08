package todo

import (
	"fmt"
	"strconv"
	"todo-console/constants"
	"todo-console/types"
)

type ItemTab struct {
	Tab
}

func (t *ItemTab) Open(c *types.RouterContext) error {
	if err := t.showItem(c.Item.GetID()); err != nil {
		return err
	}
	t.Tab.Open(c)
	return nil
}

func (t *ItemTab) HandleInput(c *types.RouterContext) *types.RouterContext {
	v, _ := strconv.Atoi(c.Input)

	switch v {
	case 0:
		c.TabName = constants.TabList
	case 1:
		// t.EditTodo()
		c.TabName = constants.TabList
	case 2:
		t.Storage.Remove(c.Item.GetID())
		c.TabName = constants.TabList
	default:
		fmt.Println("Invalid option. Please try again.")
	}

	return c
}

func (t *ItemTab) EditTodo(id string) {
	fmt.Println("Edit todo")
}

func (t *ItemTab) showItem(id int) error {
	item := t.Storage.Find(id)
	if item == nil {
		return fmt.Errorf("todo not found")
	}

	fmt.Println((*item).GetTitle())

	return nil
}

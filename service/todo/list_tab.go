package todo

import (
	"fmt"
	"strconv"
	"todo-console/constants"
	"todo-console/storage"
	"todo-console/types"
)

type ListTab struct {
	Tab
}

func (t *ListTab) Open(c *types.RouterContext) error {
	t.Tab.Open(c)
	showAllTodos(t.Storage.FindAll())
	return nil
}

func (t *ListTab) HandleInput(c *types.RouterContext) *types.RouterContext {
	v, err := strconv.Atoi(c.Input)

	if err != nil {
		fmt.Println("Invalid option. Please try again.")
		return types.NewRouterContext("list", c.Item)
	}

	if len(*t.Storage.FindAll()) > 0 && v > 0 {
		c.Item = *t.Storage.Find(v)
	}

	switch v {
	case 0:
		c.TabName = constants.TabMain
	default:
		c.TabName = constants.TabItem
	}

	return c
}

func showAllTodos(items *[]storage.Item) error {
	if len(*items) == 0 {
		fmt.Println("No todos found")
		return nil
	}

	fmt.Println("Your todos:")
	for i, item := range *items {
		fmt.Printf(
			"%d. %s - %s - %s\n",
			i+1,
			item.GetTitle(),
			item.GetDescription(),
			item.GetStatus(),
		)
	}

	return nil
}

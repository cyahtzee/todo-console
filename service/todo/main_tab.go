package todo

import (
	"fmt"
	"strconv"
	"todo-console/types"
)

type MainTab struct {
	Tab
}

func (t *MainTab) HandleInput(c *types.RouterContext) *types.RouterContext {
	v, _ := strconv.Atoi(c.Input)

	switch v {
	case 1:
		c.TabName = "list"
	case 2:
		c.TabName = "add"
	case 3:
		c.TabName = "exit"
	default:
		fmt.Println("Invalid option. Please try again.")
		c.TabName = "main"
	}

	return c
}

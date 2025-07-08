package todo

import (
	"fmt"
	"strconv"
	"todo-console/constants"
	"todo-console/types"
)

type MainTab struct {
	Tab
}

func (t *MainTab) HandleInput(c *types.RouterContext) *types.RouterContext {
	v, _ := strconv.Atoi(c.Input)

	switch v {
	case 1:
		c.TabName = constants.TabList
	case 2:
		c.TabName = constants.TabAdd
	case 3:
		c.TabName = constants.TabExit
	default:
		fmt.Println("Invalid option. Please try again.")
		c.TabName = constants.TabMain
	}

	return c
}

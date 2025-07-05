package todo

import (
	"fmt"
	"strconv"
)

type MainTab struct {
	Tab
}

func (t *MainTab) HandleInput(input string) TabInput {
	tabName := ""
	v, _ := strconv.Atoi(input)

	switch v {
	case 1:
		tabName = "list"
	case 2:
		tabName = "add"
	case 3:
		tabName = "exit"
	default:
		fmt.Println("Invalid option. Please try again.")
		tabName = "main"
	}

	return NewTabInput(tabName, nil)
}

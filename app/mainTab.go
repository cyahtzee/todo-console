package app

import (
	"fmt"
)

type MainTab struct {
	Tab
}

func (t *MainTab) HandleInput(input string) (string, string) {
	tabName := ""
	ctx := ""

	switch input {
	case "1":
		tabName = "list"
	case "2":
		tabName = "add"
	case "3":
		tabName = "exit"
	default:
		fmt.Println("Invalid option. Please try again.")
		tabName = "main"
	}

	return tabName, ctx
}

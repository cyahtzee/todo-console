package app

var tabs = map[string]string{
	"main": `
		Please select an option:
			1. List all todos
			2. Add a new todo
			3. Exit
		`,
	"list": `
		0. Back to main menu
		--------------------------------
		Select a todo to view:
		`,
	"add": `
		0. Back to main menu
		--------------------------------
		`,
	"item": `
		Select an option:
		0. Back to main menu
		--------------------------------
		1. Edit a todo
		2. Remove a todo
		`,
	"edit": `
		Select an option:
		0. Back to main menu
		--------------------------------
		`,
}

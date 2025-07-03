package app

import (
	"fmt"
)

type ItemTab struct {
	Tab
}

func (t *ItemTab) Open() error {
	t.Tab.Open()
	if err := t.showToDo(); err != nil {
		return err
	}
	return nil
}

func (t *ItemTab) HandleInput(input string) (string, string) {
	tabName := ""
	ctx := ""

	switch input {
	case "0":
		tabName = "list"
	case "1":
		// t.EditTodo()
		tabName = "list"
	case "2":
		t.RemoveTodo()
		tabName = "list"
	default:
		fmt.Println("Invalid option. Please try again.")
	}

	return tabName, ctx
}

func (t *ItemTab) EditTodo(id string) {
	fmt.Println("Edit todo")
}

func (t *ItemTab) RemoveTodo() {
	for i, item := range *t.Items {
		fmt.Println("Checking todo: ", item.GetID())
		fmt.Println("Checking input: ", t.Ctx)
		if fmt.Sprintf("%d", item.GetID()) == t.Ctx {
			fmt.Println("Removing todo: ", item.GetTitle())
			*t.Items = append((*t.Items)[:i], (*t.Items)[i+1:]...)
			fmt.Println("Todo removed")
			return
		}
	}
	fmt.Println("Todo not found")
}

func (t *ItemTab) showToDo() error {
	for _, item := range *t.Items {
		fmt.Println("Checking todo: ", fmt.Sprintf("%d", item.GetID()))
		fmt.Println("Checking input: ", t.GetCtx())
		fmt.Println("Checking if input is a number: ", t.GetCtx() == fmt.Sprintf("%d", item.GetID()))
		if fmt.Sprintf("%d", item.GetID()) == t.GetCtx() {
			fmt.Println(item.GetTitle())
			return nil
		}
	}

	return fmt.Errorf("todo not found")
}

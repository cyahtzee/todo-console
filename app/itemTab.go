package app

import (
	"fmt"
	"strconv"
)

type ItemTab struct {
	Tab
}

func (t *ItemTab) Open() error {
	if err := t.showToDo(); err != nil {
		return err
	}
	t.Tab.Open()
	return nil
}

func (t *ItemTab) HandleInput(input string) TabInput {
	tabName := ""
	ctx := t.Ctx
	v, _ := strconv.Atoi(input)

	switch v {
	case 0:
		tabName = "list"
	case 1:
		// t.EditTodo()
		tabName = "list"
	case 2:
		t.RemoveTodo()
		tabName = "list"
	default:
		fmt.Println("Invalid option. Please try again.")
	}

	return NewTabInput(tabName, ctx)
}

func (t *ItemTab) EditTodo(id string) {
	fmt.Println("Edit todo")
}

func (t *ItemTab) RemoveTodo() {
	for i, item := range *t.Items {
		if item == *t.Ctx {
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
		if item == *t.Ctx {
			fmt.Println(item.GetTitle())
			return nil
		}
	}

	return fmt.Errorf("todo not found")
}

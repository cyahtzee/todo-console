package todo

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
	for _, item := range *t.Storage.FindAll() {
		if item.GetID() == (*t.Ctx).GetID() {
			fmt.Println("Removing todo: ", item.GetTitle())
			t.Storage.Remove(item.GetID())
			fmt.Println("Todo removed")
			return
		}
	}
	fmt.Println("Todo not found")
}

func (t *ItemTab) showToDo() error {
	item := t.Storage.Find((*t.Ctx).GetID())
	if item == nil {
		return fmt.Errorf("todo not found")
	}

	fmt.Println((*item).GetTitle())
	return nil
}

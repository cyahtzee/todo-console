package app

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

func NewTodo(id int, title string, description string, completed bool) *Todo {
	return &Todo{ID: id, Title: title, Description: description, Completed: completed}
}

func (t *Todo) GetID() int {
	return t.ID
}

func (t *Todo) GetTitle() string {
	return t.Title
}

func (t *Todo) GetDescription() string {
	return t.Description
}

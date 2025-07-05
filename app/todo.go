package app

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   string
}

var requiredFields = []string{
	"title",
	"description",
	"completed",
}

func NewTodo(id int, title string, description string, completed string) *Todo {
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

func (t *Todo) SetField(field string, value string) {
	switch field {
	case "title":
		t.Title = value
	case "description":
		t.Description = value
	case "completed":
		t.Completed = value
	}
}

func (t *Todo) Valid() bool {
	return t.Title != "" && t.Description != "" && t.Completed != ""
}

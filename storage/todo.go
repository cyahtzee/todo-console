package storage

import (
	"fmt"
	"todo-console/constants"
)

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   constants.Status
}

func NewTodo(id int, title string, description string) *Todo {
	return &Todo{ID: id, Title: title, Description: description, Completed: constants.PENDING}
}

func (t *Todo) GetID() int {
	return t.ID
}

func (t *Todo) SetID(id int) {
	t.ID = id
}

func (t *Todo) GetTitle() string {
	return t.Title
}

func (t *Todo) GetDescription() string {
	return t.Description
}

func (t *Todo) GetStatus() string {
	return t.Completed.String()
}

func (t *Todo) SetField(field string, value string) error {
	switch field {
	case "title":
		t.Title = value
	case "description":
		t.Description = value
	case "completed":
		if !constants.Status(value).Validate() {
			return fmt.Errorf("invalid completed value: %s", value)
		}
		t.Completed = constants.Status(value)
	}
	return nil
}

func (t *Todo) GetField(field string) string {
	fieldValue := ""
	switch field {
	case "title":
		fieldValue = t.Title
	case "description":
		fieldValue = t.Description
	case "completed":
		fieldValue = t.Completed.String()
	}
	return fieldValue
}

func (t *Todo) Validate(fields []string) bool {
	for _, field := range fields {
		if t.GetField(field) == "" {
			return false
		}
	}
	return true
}

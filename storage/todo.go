package storage

type Todo struct {
	ID          int
	Title       string
	Description string
	Completed   string
}

func NewTodo(id int, title string, description string, completed string) *Todo {
	return &Todo{ID: id, Title: title, Description: description, Completed: completed}
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

func (t *Todo) GetField(field string) string {
	fieldValue := ""
	switch field {
	case "title":
		fieldValue = t.Title
	case "description":
		fieldValue = t.Description
	case "completed":
		fieldValue = t.Completed
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

package storage

type Storage struct {
	Items *[]Item
}

type Item interface {
	GetID() int
	SetID(id int)
	GetTitle() string
	GetDescription() string
	GetStatus() string
	SetField(field string, value string) error
	GetField(field string) string
	Validate(fields []string) bool
}

type Store interface {
	Create(todo *Item)
	Find(id int) *Item
	Remove(id int)
	FindAll() *[]Item
}

func NewStorage(items *[]Item) *Storage {
	return &Storage{Items: items}
}

func (s *Storage) Create(item Item) {
	id := len(*s.Items) + 1
	item.SetID(id)
	*s.Items = append(*s.Items, item)
}

func (s *Storage) Remove(id int) {
	for i, item := range *s.Items {
		if item.GetID() == id {
			*s.Items = append((*s.Items)[:i], (*s.Items)[i+1:]...)
		}
	}
}

func (s *Storage) Find(id int) *Item {
	for _, item := range *s.Items {
		if item.GetID() == id {
			return &item
		}
	}
	return nil
}

func (s *Storage) FindAll() *[]Item {
	return s.Items
}

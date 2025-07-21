package todo

import (
	"fmt"
	"todo-console/storage"
	"todo-console/types"
)

type Tab struct {
	Active  bool
	Name    string
	Step    int
	Storage storage.Store
}

type TabInterface interface {
	HandleInput(c *types.RouterContext) *types.RouterContext
	Open(c *types.RouterContext) error
	GetUI() string
	GetName() string
	GetStatus() bool
	SetActive()
	Close() error
}

func (t *Tab) GetUI() string {
	return tabs[t.Name]
}

func (t *Tab) GetName() string {
	return t.Name
}

func (t *Tab) GetStatus() bool {
	return t.Active
}

func (t *Tab) Open(c *types.RouterContext) error {
	t.SetActive()
	fmt.Println(t.GetUI())
	return nil
}

func (t *Tab) SetActive() {
	t.Active = true
}

func (t *Tab) Close() error {
	t.Active = false
	return nil
}

func (t *Tab) GetProperty() string {
	return requiredFields[t.Step]
}

func (t *Tab) NextStep() {
	t.Step++
}

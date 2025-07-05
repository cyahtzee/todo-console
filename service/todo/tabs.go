package todo

import (
	"fmt"
	"todo-console/storage"
)

type Tab struct {
	Active  bool
	Name    string
	Step    int
	Ctx     *storage.Item
	Storage *storage.Storage
}

type TabInterface interface {
	HandleInput(input string) TabInput
	GetUI() string
	GetName() string
	GetStatus() bool
	Open() error
	SetActive()
	Close() error
	GetCtx() *storage.Item
	SetCtx(ctx *storage.Item)
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

func (t *Tab) Open() error {
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

func (t *Tab) GetCtx() *storage.Item {
	return t.Ctx
}

func (t *Tab) SetCtx(ctx *storage.Item) {
	t.Ctx = ctx
}

func (t *Tab) GetProperty() string {
	return requiredFields[t.Step]
}

func (t *Tab) NextStep() {
	t.Step++
}

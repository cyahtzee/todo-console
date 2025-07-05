package app

import (
	"fmt"
)

type Tab struct {
	Active bool
	Name   string
	Step   int
	Ctx    *Todo
	Items  *[]Todo
}

type TabInterface interface {
	HandleInput(input string) TabInput
	GetUI() string
	GetName() string
	GetStatus() bool
	Open() error
	SetActive()
	Close() error
	GetCtx() *Todo
	SetCtx(ctx *Todo)
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

func (t *Tab) GetCtx() *Todo {
	return t.Ctx
}

func (t *Tab) SetCtx(ctx *Todo) {
	t.Ctx = ctx
}

func (t *Tab) GetProperty() string {
	return requiredFields[t.Step]
}

func (t *Tab) NextStep() {
	t.Step++
}

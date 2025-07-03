package app

import (
	"fmt"
)

type Tab struct {
	Active bool
	Name   string
	Ctx    string
	Items  *[]Todo
}

type TabInterface interface {
	HandleInput(input string) (string, string)
	GetUI() string
	GetName() string
	GetStatus() bool
	Open() error
	SetActive()
	Close() error
	GetCtx() string
	SetCtx(ctx string)
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
	fmt.Println(t.GetUI())
	t.SetActive()
	return nil
}

func (t *Tab) SetActive() {
	t.Active = true
}

func (t *Tab) Close() error {
	t.Active = false
	return nil
}

func (t *Tab) GetCtx() string {
	return t.Ctx
}

func (t *Tab) SetCtx(ctx string) {
	t.Ctx = ctx
}

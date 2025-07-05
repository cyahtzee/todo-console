package todo

import "todo-console/storage"

type TabInput struct {
	Ctx     *storage.Item
	TabName string
}

func NewTabInput(tabName string, ctx *storage.Item) TabInput {
	if ctx == nil {
		todo := storage.NewTodo(0, "", "", "")
		var item storage.Item = todo
		ctx = &item
	}

	return TabInput{
		TabName: tabName,
		Ctx:     ctx,
	}
}

func (t *TabInput) GetCtx() *storage.Item {
	return t.Ctx
}

func (t *TabInput) SetCtx(ctx *storage.Item) {
	t.Ctx = ctx
}

func (t *TabInput) GetTabName() string {
	return t.TabName
}

func (t *TabInput) SetTabName(tabName string) {
	t.TabName = tabName
}

package app

type TabInput struct {
	Ctx     *Todo
	TabName string
}

func NewTabInput(tabName string, ctx *Todo) TabInput {
	if ctx == nil {
		ctx = &Todo{}
	}

	return TabInput{
		TabName: tabName,
		Ctx:     ctx,
	}
}

func (t *TabInput) GetCtx() *Todo {
	return t.Ctx
}

func (t *TabInput) SetCtx(ctx *Todo) {
	t.Ctx = ctx
}

func (t *TabInput) GetTabName() string {
	return t.TabName
}

func (t *TabInput) SetTabName(tabName string) {
	t.TabName = tabName
}

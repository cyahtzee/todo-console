package app

type TabInput struct {
	Ctx     string
	TabName string
}

func (t *TabInput) GetCtx() string {
	return t.Ctx
}

func (t *TabInput) GetTabName() string {
	return t.TabName
}

func (t *TabInput) SetCtx(ctx string) {
	t.Ctx = ctx
}

func (t *TabInput) SetTabName(tabName string) {
	t.TabName = tabName
}

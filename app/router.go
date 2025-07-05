package app

import (
	"bufio"
	"os"
	"strings"
	"todo-console/service/todo"
	"todo-console/types"
)

type Router struct {
	Reader      *bufio.Reader
	Tab         *todo.TabInterface
	PreviousTab *todo.TabInterface
	TabsChannel chan string
	Ctx         *types.RouterContext
}

func NewRouter(ctx *types.RouterContext) *Router {
	return &Router{
		Reader:      bufio.NewReader(os.Stdin),
		TabsChannel: make(chan string, 10),
		Ctx:         ctx,
	}
}

func (r *Router) HandleInput() error {
	input, err := r.Reader.ReadString('\n')
	if err != nil {
		return err
	}
	r.Ctx.Input = strings.TrimSpace(input)
	ctx := (*r.Tab).HandleInput(r.Ctx)
	r.SetCtx(ctx)

	r.TabsChannel <- ctx.TabName
	return nil
}

func (r *Router) SetCtx(ctx *types.RouterContext) {
	r.Ctx = ctx
}

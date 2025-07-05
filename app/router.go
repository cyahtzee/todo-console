package app

import (
	"bufio"
	"os"
	"strings"
	"todo-console/service/todo"
)

type Router struct {
	Reader      *bufio.Reader
	Tab         *todo.TabInterface
	PreviousTab *todo.TabInterface
	TabsChannel chan todo.TabInput
}

func NewRouter() *Router {
	return &Router{
		Reader:      bufio.NewReader(os.Stdin),
		TabsChannel: make(chan todo.TabInput, 10),
	}
}

func (r *Router) HandleInput() error {
	input, err := r.Reader.ReadString('\n')
	if err != nil {
		return err
	}
	input = strings.TrimSpace(input)

	msg := (*r.Tab).HandleInput(input)

	r.TabsChannel <- msg
	return nil
}

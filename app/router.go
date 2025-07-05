package app

import (
	"bufio"
	"strings"
)

type Router struct {
	Reader      *bufio.Reader
	Tab         *TabInterface
	PreviousTab *TabInterface
	TabsChannel chan TabInput
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

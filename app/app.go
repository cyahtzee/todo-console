package app

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type App struct {
	Running bool
	Items   *[]Todo
	UI      *[]TabInterface
	wg      sync.WaitGroup
	router  *Router
}

func NewApp(items *[]Todo, ui *[]TabInterface) *App {
	return &App{
		Running: true,
		Items:   items,
		UI:      ui,
		router: &Router{
			Reader:      bufio.NewReader(os.Stdin),
			TabsChannel: make(chan TabInput, 10),
		},
	}
}

func (a *App) Run() error {
	a.Running = true

	go a.handleTabMessages()
	a.SetCurrentTab()

	if a.router.Tab == nil {
		return fmt.Errorf("no active tab found")
	}

	for a.Running {
		a.PrintHeader()

		if err := (*a.router.Tab).Open(); err != nil {
			fmt.Println("Error opening tab: ", err)
			a.router.TabsChannel <- TabInput{TabName: "main", Ctx: ""}
		}

		a.wg.Add(1)
		a.router.HandleInput()
		a.wg.Wait()
	}

	close(a.router.TabsChannel)

	return nil
}

func (a *App) handleTabMessages() {
	for a.Running {
		msg := <-a.router.TabsChannel
		switch msg.TabName {
		case "exit":
			a.Stop()
		default:
			a.router.PreviousTab = a.router.Tab
			a.SwitchTab(msg)
		}
		a.wg.Done()
	}
}

func (a *App) Stop() error {
	a.Running = false
	return nil
}

func (a *App) PrintHeader() {
	fmt.Println("Welcome to the Todo Console")
}

func (a *App) GetCurrentTab() *TabInterface {
	for _, tab := range *a.UI {
		if tab.GetStatus() {
			return &tab
		}
	}
	return nil
}

func (a *App) SwitchTab(msg TabInput) error {
	fmt.Println("Switching to tab: ", msg.TabName)
	for _, tab := range *a.UI {
		tab.Close()
	}

	for _, tab := range *a.UI {
		if tab.GetName() == msg.TabName {
			tab.SetActive()
			tab.SetCtx(msg.Ctx)
			a.router.Tab = &tab
			break
		}
	}
	return nil
}

func (a *App) SetCurrentTab() {
	a.router.Tab = a.GetCurrentTab()
}

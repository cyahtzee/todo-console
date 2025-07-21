package app

import (
	"fmt"
	"sync"
	"todo-console/constants"
	"todo-console/service/todo"
	"todo-console/storage"
)

type App struct {
	Running     bool
	Storage     storage.Store
	Controllers *[]todo.TabInterface
	wg          sync.WaitGroup
	Router      *Router
}

func NewApp(storage storage.Store, router *Router, controllers *[]todo.TabInterface) *App {
	return &App{
		Running:     false,
		Storage:     storage,
		Controllers: controllers,
		Router:      router,
	}
}

func (a *App) Run() error {
	a.Running = true
	a.Router.Tab = a.GetCurrentTab()

	go a.handleTabMessages()

	if a.Router.Tab == nil {
		return fmt.Errorf("no active tab found")
	}

	for a.Running {
		fmt.Println("Welcome to the Todo Console")

		if err := (*a.Router.Tab).Open(a.Router.Ctx); err != nil {
			fmt.Println("Error opening tab: ", err)
			fmt.Println("Fallback to previous tab: ", (*a.Router.PreviousTab).GetName())
			a.Router.Tab = a.Router.PreviousTab
			continue
		}

		a.wg.Add(1)
		a.Router.HandleInput()
		a.wg.Wait()
	}

	close(a.Router.TabsChannel)

	return nil
}

func (a *App) handleTabMessages() {
	for a.Running {
		msg := <-a.Router.TabsChannel
		switch msg {
		case constants.TabExit:
			a.Stop()
		default:
			a.Router.PreviousTab = a.Router.Tab
			a.SwitchTab(msg)
		}
		a.wg.Done()
	}
}

func (a *App) Stop() error {
	a.Running = false
	return nil
}

func (a *App) SwitchTab(msg string) error {
	for _, tab := range *a.Controllers {
		tab.Close()
	}

	for _, tab := range *a.Controllers {
		if tab.GetName() == msg {
			tab.SetActive()
			a.Router.Tab = &tab
			break
		}
	}
	return nil
}

func (a *App) GetCurrentTab() *todo.TabInterface {
	for _, tab := range *a.Controllers {
		if tab.GetStatus() {
			return &tab
		}
	}
	return nil
}

func (a *App) SetCurrentTab() {
	a.Router.Tab = a.GetCurrentTab()
}

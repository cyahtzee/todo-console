package app

import (
	"fmt"
	"sync"
	"todo-console/service/todo"
	"todo-console/storage"
)

type App struct {
	Running     bool
	Storage     *storage.Storage
	Controllers *[]todo.TabInterface
	wg          sync.WaitGroup
	router      *Router
}

func NewApp(storage *storage.Storage, controllers *[]todo.TabInterface, router *Router) *App {
	return &App{
		Running:     true,
		Storage:     storage,
		Controllers: controllers,
		router:      router,
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
		fmt.Println("Welcome to the Todo Console")

		if err := (*a.router.Tab).Open(); err != nil {
			fmt.Println("Error opening tab: ", err)
			fmt.Println("Fallback to previous tab: ", (*a.router.PreviousTab).GetName())
			a.router.Tab = a.router.PreviousTab
			continue
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

func (a *App) SwitchTab(msg todo.TabInput) error {
	for _, tab := range *a.Controllers {
		tab.Close()
	}

	for _, tab := range *a.Controllers {
		if tab.GetName() == msg.TabName {
			tab.SetActive()
			tab.SetCtx(msg.Ctx)
			a.router.Tab = &tab
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
	a.router.Tab = a.GetCurrentTab()
}

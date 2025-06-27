package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Menu struct {
	Running bool
	Items   *[]Todo
}

func (c *Menu) Run() error {
	c.Running = true
	for c.Running {
		fmt.Println(`
		Welcome to the Todo Console
		Please select an option:
			1. List all todos
			2. Add a new todo
			3. Exit
		`)
		input := ""
		fmt.Scanln(&input)

		switch input {
		case "1":
			c.List()
		case "2":
			c.Add()
		case "3":
			return c.Stop()
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
	return nil
}

func (c *Menu) Stop() error {
	c.Running = false
	return nil
}

func (c *Menu) List() error {
	for _, item := range *c.Items {
		fmt.Printf("%d. %s\n", item.GetID(), item.GetTitle())
	}
	return nil
}

func (c *Menu) Add() error {
	fmt.Println("Enter the todo:")
	reader := bufio.NewReader(os.Stdin)
	todo, err := reader.ReadString('\n')

	if err != nil {
		return err
	}

	todo = strings.TrimSpace(todo)

	if todo == "" {
		fmt.Println("Todo is empty")
		return nil
	}

	*c.Items = append(*c.Items, *NewTodo(len(*c.Items)+1, todo, "", false))
	fmt.Printf("Added: %s\n", todo)

	return nil
}

func (c *Menu) Get(id int) error {
	for _, item := range *c.Items {
		if item.GetID() == id {
			fmt.Printf("Todo: %s\n", item.GetTitle())
		}
	}

	return nil
}

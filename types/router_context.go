package types

import (
	"todo-console/storage"
)

type RouterContext struct {
	TabName string
	Input   string
	Item    storage.Item
}

func NewRouterContext(tabName string, item storage.Item) *RouterContext {
	return &RouterContext{
		TabName: tabName,
		Item:    item,
	}
}

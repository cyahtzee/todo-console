package types_test

import (
	"testing"
	"todo-console/storage"
	"todo-console/types"
)

func TestRouterContext_WithDifferentTabNames(t *testing.T) {
	tabNames := []string{"main", "add", "edit", "list", "item", "delete", "search"}

	for _, tabName := range tabNames {
		t.Run("tab name: "+tabName, func(t *testing.T) {
			item := storage.Todo{}
			context := types.NewRouterContext(tabName, &item)

			if context == nil {
				t.Fatal("Expected RouterContext to be created, got nil")
			}

			if context.TabName != tabName {
				t.Errorf("Expected TabName to be '%s', got '%s'", tabName, context.TabName)
			}
		})
	}
}

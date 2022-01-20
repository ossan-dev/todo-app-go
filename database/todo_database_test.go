package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/error_handler"
	"todo-app-go.com/v1/model"
)

func TestGetById(t *testing.T) {
	store := database.NewStubTodoStore(
		&map[int]model.Todo{
			1: model.NewTodo(1, "FirstTodo", false),
			2: model.NewTodo(2, "SecondTodo", false),
		},
	)

	t.Run("todo present in collection", func(t *testing.T) {
		got, err := store.GetTodoById(1)

		assert.Equal(t, "FirstTodo", got)
		assert.NoError(t, err)
	})

	t.Run("todo not present in collection", func(t *testing.T) {
		got, err := store.GetTodoById(4)

		assert.Equal(t, "", got)
		assert.IsType(t, error_handler.ErrNotFound, err)
	})
}

func TestAdd(t *testing.T) {
	store := database.NewStubTodoStore(
		&map[int]model.Todo{},
	)

	t.Run("add todo when it's correct", func(t *testing.T) {
		got, err := store.AddTodo("Example todo")
		assert.Equal(t, 1, got)
		assert.NoError(t, err)
	})

	t.Run("return an error with blank description", func(t *testing.T) {
		got, err := store.AddTodo("")
		assert.Equal(t, 0, got)
		assert.Error(t, err)
	})
}

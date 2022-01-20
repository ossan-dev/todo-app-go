package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/error_handler"
	"todo-app-go.com/v1/model"
	"todo-app-go.com/v1/util"
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

		util.AssertResponseBody(t, got, "FirstTodo")
		util.AssertNoError(t, err)
	})

	t.Run("todo not present in collection", func(t *testing.T) {
		got, err := store.GetTodoById(4)

		util.AssertResponseBody(t, got, "")
		util.AssertError(t, err, error_handler.ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	store := database.NewStubTodoStore(
		&map[int]model.Todo{},
	)

	got := store.AddTodo("Example todo")

	assert.Equal(t, 1, got)
}

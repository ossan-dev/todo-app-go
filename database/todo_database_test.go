package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"todo-app-go.com/v1/database"
	"todo-app-go.com/v1/error_handler"
	"todo-app-go.com/v1/model"
)

func TestGetById(t *testing.T) {
	todo := model.NewTodo(1, "FirstTodo", false)
	store := database.NewStubTodoStore(
		&map[int]model.Todo{
			1: todo,
		},
	)

	t.Run("todo present in collection", func(t *testing.T) {
		got, err := store.GetTodoById(1)

		assert.Equal(t, &todo, got)
		assert.NoError(t, err)
	})

	t.Run("todo not present in collection", func(t *testing.T) {
		got, err := store.GetTodoById(4)

		assert.Nil(t, got)
		assert.IsType(t, error_handler.ErrNotFound, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("GetAllTodos should return all the todos added", func(t *testing.T) {
		todos := []model.Todo{
			model.NewTodo(1, "FirstTodo", false),
			model.NewTodo(2, "SecondTodo", true),
			model.NewTodo(3, "ThirdTodo", false),
		}
		store := database.NewStubTodoStore(
			&map[int]model.Todo{
				1: todos[0],
				2: todos[1],
				3: todos[2],
			},
		)

		// TODO: make other assertions
		// todos is unsorted
		assert.Equal(t, len(store.GetAllTodos()), len(todos))
	})

	t.Run("GetAllTodos should return [] when no todos are present", func(t *testing.T) {
		todos := make([]model.Todo, 0)
		store := database.NewStubTodoStore(
			&map[int]model.Todo{},
		)

		assert.Equal(t, store.GetAllTodos(), todos)
		assert.Equal(t, len(store.GetAllTodos()), len(todos))
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

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
		todos := map[int]model.Todo{
			1: model.NewTodo(1, "FirstTodo", false),
			2: model.NewTodo(2, "SecondTodo", true),
			3: model.NewTodo(3, "ThirdTodo", false),
		}
		store := database.NewStubTodoStore(
			&todos,
		)

		assert.Equal(t, store.GetAllTodos(), todos)
	})

	// t.Run("GetAllTodos should return [] when no todos are present", func(t *testing.T) {
	// 	todos := []models.Todo{}

	// 	todoManager := &TodoManager{todos}

	// 	util.AssertCollectionsEqual(t, todoManager.GetAllTodos(), todos)
	// })
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

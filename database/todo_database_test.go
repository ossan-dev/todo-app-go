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

		got := store.GetAllTodos()
		util.SortTodoSliceById(todos)
		util.SortTodoSliceById(got)

		assert.Equal(t, todos, got)
	})

	t.Run("GetAllTodos should return [] when no todos are present", func(t *testing.T) {
		todos := make([]model.Todo, 0)
		store := database.NewStubTodoStore(
			&map[int]model.Todo{},
		)

		assert.Equal(t, store.GetAllTodos(), todos)
	})
}

func TestGetByStatus(t *testing.T) {
	t.Run("get completed todos return non-empty collections", func(t *testing.T) {
		todos := []model.Todo{
			model.NewTodo(1, "FirstTodo", false),
			model.NewTodo(2, "SecondTodo", true),
			model.NewTodo(3, "ThirdTodo", true),
		}
		store := database.NewStubTodoStore(
			&map[int]model.Todo{
				1: todos[0],
				2: todos[1],
				3: todos[2],
			},
		)

		got := store.GetByStatus(true)
		util.SortTodoSliceById(got)
		assert.Equal(t, []model.Todo{todos[1], todos[2]}, got)
	})

	t.Run("get completed todos return empty collections", func(t *testing.T) {
		todos := []model.Todo{
			model.NewTodo(1, "FirstTodo", false),
			model.NewTodo(2, "SecondTodo", false),
		}
		store := database.NewStubTodoStore(
			&map[int]model.Todo{
				1: todos[0],
				2: todos[1],
			},
		)

		got := store.GetByStatus(true)
		util.SortTodoSliceById(got)
		assert.Equal(t, []model.Todo{}, got)
	})
}

func TestAdd(t *testing.T) {
	store := database.NewStubTodoStore(
		&map[int]model.Todo{},
	)

	t.Run("add todo when it's correct", func(t *testing.T) {
		got, err := store.AddTodo(model.NewTodo(1, "Example todo", false))
		assert.Equal(t, 1, got)
		assert.NoError(t, err)
	})

	t.Run("return an error with blank description", func(t *testing.T) {
		got, err := store.AddTodo(model.NewTodo(1, "", true))
		assert.Equal(t, 0, got)
		assert.IsType(t, error_handler.ErrTodoNotValid, err)
	})
}

func TestUpdate(t *testing.T) {
	todos := []model.Todo{
		model.NewTodo(1, "FirstTodo", false),
	}
	store := database.NewStubTodoStore(
		&map[int]model.Todo{
			1: todos[0],
		},
	)

	t.Run("update existing todo returns updated id", func(t *testing.T) {
		todo := model.NewTodo(1, "UpdatedTodo", true)
		got, err := store.UpdateTodo(todo)

		assert.Equal(t, todo.Id, *got)
		assert.Nil(t, err)
	})

	t.Run("update not existing todo returns err", func(t *testing.T) {
		todo := model.NewTodo(4, "UpdatedTodo", true)
		got, err := store.UpdateTodo(todo)

		assert.IsType(t, error_handler.ErrNotFound, err)
		assert.Nil(t, got)
	})
}

// func TestDelete(t *testing.T) {
// 	todos := []models.Todo{
// 		models.NewTodo(1, "FirstTodo", false),
// 		models.NewTodo(2, "SecondTodo", false),
// 		models.NewTodo(3, "ThirdTodo", true),
// 	}

// 	todoManager := &TodoManager{todos}

// 	t.Run("delete todo", func(t *testing.T) {
// 		todoManager.DeleteById(2)
// 		_, err := todoManager.GetById(2)
// 		utils.AssertError(t, err, ErrTodoNotFound)
// 	})

// 	t.Run("not delete todo if not exists", func(t *testing.T) {
// 		err := todoManager.DeleteById(4)
// 		utils.AssertError(t, err, ErrTodoNotFound)
// 	})
// }

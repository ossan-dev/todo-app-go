package crud

import (
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestGetById(t *testing.T) {
	todos := []models.Todo{
		models.NewTodo(1, "FirstTodo", false),
		models.NewTodo(2, "SecondTodo", false),
		models.NewTodo(3, "ThirdTodo", true),
	}

	todoManager := &TodoManager{todos}

	t.Run("todo present in collection", func(t *testing.T) {
		got, err := todoManager.GetById(1)

		utils.AssertTodosEqual(t, got, &todos[0])
		utils.AssertNoError(t, err)
	})

	t.Run("todo not present in collection", func(t *testing.T) {
		got, err := todoManager.GetById(4)

		utils.AssertTodosEqual(t, got, nil)
		utils.AssertError(t, err, ErrTodoNotFound)
	})
}

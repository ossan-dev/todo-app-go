package crud

import (
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestGetById(t *testing.T) {
	todos := []models.Todo{
		{Id: 1, Description: "FirstTodo", IsCompleted: false},
		{Id: 2, Description: "SecondTodo", IsCompleted: false},
		{Id: 3, Description: "ThirdTodo", IsCompleted: false},
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

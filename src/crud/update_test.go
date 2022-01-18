package crud

import (
	"testing"

	"todo-app-go.com/v1/src/models"
	"todo-app-go.com/v1/src/utils"
)

func TestUpdate(t *testing.T) {
	todos := []models.Todo{
		models.NewTodo(1, "FirstTodo", false),
		models.NewTodo(2, "SecondTodo", false),
		models.NewTodo(3, "ThirdTodo", true),
	}

	todoManager := &TodoManager{todos}

	t.Run("update existing todo", func(t *testing.T) {
		todoUpdated := models.NewTodo(2, "Updated text", false)
		todoManager.Update(todoUpdated)

		got, err := todoManager.GetById(todoUpdated.Id)

		utils.AssertNoError(t, err)
		utils.AssertTodosEqual(t, got, &todoUpdated)
	})

	t.Run("update not existing todo", func(t *testing.T) {
		todoUpdated := models.NewTodo(4, "Updated text", false)
		err := todoManager.Update(todoUpdated)

		utils.AssertError(t, err, ErrTodoNotFound)
	})
}
